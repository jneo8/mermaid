package mermaid

import (
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"go.uber.org/dig"
	"strings"
)

// Mermaid .
type Mermaid struct {
	Container *dig.Container
	CMD       *cobra.Command
	Config    *viper.Viper
	Logger    *log.Logger
	ENVPrefix string
}

// NewMermaid return new mermaid worker.
func NewMermaid(
	cmd *cobra.Command,
	config *viper.Viper,
	logger *log.Logger,
	envPrefix string,
) (*Mermaid, error) {
	worker := Mermaid{
		Container: dig.New(),
		CMD:       cmd,
		Config:    config,
		Logger:    logger,
		ENVPrefix: envPrefix,
	}
	if err := worker.Bind(); err != nil {
		return nil, err
	}
	return &worker, nil
}

// Run .
func Run(cmd *cobra.Command, runable interface{}, initializers ...interface{}) error {
	mermaid, err := NewMermaid(
		cmd,
		viper.New(),
		NewLogger(),
		cmd.Use,
	)
	if err != nil {
		return err
	}
	return mermaid.Execute(runable, initializers...)
}

// Execute .
func (m *Mermaid) Execute(runable interface{}, initializers ...interface{}) error {
	container := m.Container

	// Provide config && logger by default.
	initializers = append(initializers, func() *viper.Viper { return m.Config })
	initializers = append(initializers, func() *log.Logger { return m.Logger })
	for _, initFn := range initializers {
		if err := container.Provide(initFn); err != nil {
			m.Logger.Error(err)
			return err
		}
	}

	return dig.RootCause(container.Invoke(runable))
}

// Bind .
func (m *Mermaid) Bind() error {
	m.BindViper()
	err := BindFlags(m.CMD, m.Config)
	if err != nil {
		m.Logger.Error(err)
		return err
	}
	if err := BindContainer(m.Config, m.Container); err != nil {
		m.Logger.Error(err)
		return err
	}
	return nil
}

// BindViper will
//  - load config file if exists.
//  - load environment with prefix.
func (m *Mermaid) BindViper() {

	// Read config from giving file path or filename.yaml.
	var cfgFile string
	if cfgFlag := m.CMD.Flags().Lookup("config"); cfgFlag != nil {
		cfgFile = cfgFlag.Value.String()
	}
	m.Config.SetDefault("config", cfgFile)

	if cfgFile != "" {
		m.Config.SetConfigFile(cfgFile)
		m.Logger.Infof("Use config %v", cfgFile)
	} else {
		var cfgName string
		if cfgFlag := m.CMD.Flags().Lookup("config_name"); cfgFlag != nil {
			cfgName = cfgFlag.Value.String()
		}
		m.Config.SetDefault("config_name", cfgName)
		m.Config.SetConfigName(cfgName)
		m.Config.SetConfigType("yaml")
		m.Config.AddConfigPath(".")
		m.Config.AddConfigPath("/")
		m.Config.AddConfigPath("$HOME")
		m.Config.AddConfigPath("./config")
		m.Logger.Infof("Use config name %v", cfgName)
	}

	if err := m.Config.ReadInConfig(); err != nil {
		m.Logger.Warning(err)
	} else {
		m.Logger.Infof("Using config file %v", m.Config.ConfigFileUsed())
	}

	// Load var from env.
	m.Config.AutomaticEnv()
	m.Config.SetEnvPrefix(m.ENVPrefix)
	m.Config.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
}
