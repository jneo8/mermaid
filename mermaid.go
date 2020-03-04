package mermaid

import (
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"go.uber.org/dig"
	"strings"
)

// Worker .
type Worker struct {
	Container *dig.Container
	CMD       *cobra.Command
	Config    *viper.Viper
	Logger    *log.Logger
	ENVPrefix string
}

// New return new worker .
func New(
	cmd *cobra.Command,
	config *viper.Viper,
	logger *log.Logger,
	envPrefix string,
) *Worker {
	worker := Worker{
		Container: dig.New(),
		CMD:       cmd,
		Config:    config,
		Logger:    logger,
		ENVPrefix: envPrefix,
	}
	worker.Bind()
	return &worker
}

// Run .
func (w *Worker) Run(runable interface{}, initializers ...interface{}) error {
	container := w.Container
	for _, initFn := range initializers {
		if err := container.Provide(initFn); err != nil {
			w.Logger.Error(err)
			return err
		}
	}

	if err := container.Provide(
		func() *viper.Viper { return w.Config },
		dig.Name("config"),
	); err != nil {
		return err
	}

	if err := container.Provide(
		func() *log.Logger { return w.Logger },
		dig.Name("logger"),
	); err != nil {
		return err
	}

	return dig.RootCause(container.Invoke(runable))
}

// Bind .
func (w *Worker) Bind() {
	w.BindViper()
	if err := BindContainer(w.Config, w.Container); err != nil {
		w.Logger.Error(err)
	}
}

// BindViper will
//  - load config file if exists.
//  - load environment with prefix.
//  - bind cmd flag to cfg file and provide to container.
func (w *Worker) BindViper() {
	// Read config from giving file path or filename.yaml.
	cfgFile := w.Config.GetString("config")
	if cfgFile != "" {
		w.Config.SetConfigFile(cfgFile)
		w.Logger.Infof("Use config %v", cfgFile)
	} else {
		w.Config.SetDefault("config_name", "default.yaml")
		cfgName := w.Config.GetString("config_name")
		w.Config.SetConfigName(cfgName)
		w.Config.SetConfigType("yaml")
		w.Config.AddConfigPath(".")
		w.Config.AddConfigPath("./config")
		w.Logger.Infof("Use config name %v", cfgName)
	}

	if err := w.Config.ReadInConfig(); err != nil {
		w.Logger.Warning(err)
	} else {
		w.Logger.Infof("Using config file %v", w.Config.ConfigFileUsed())
	}

	// Load var from env.
	w.Config.AutomaticEnv()
	w.Config.SetEnvPrefix(w.ENVPrefix)
	w.Config.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	w.Config.SetTypeByDefaultValue(true)

	err := BindFlags(w.CMD, w.Config)
	if err != nil {
		w.Logger.Error(err)
	}
}
