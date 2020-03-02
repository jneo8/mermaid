package mermaid

import (
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"go.uber.org/dig"
	"strings"
)

type Worker struct {
	Container *dig.Container
	CMD       *cobra.Command
	Config    *viper.Viper
	Logger    *log.Logger
	ENVPrefix string
}

func New(
	envPrefix string,
	c *dig.Container,
	config *viper.Viper,
	logger *log.Logger,
	cmd *cobra.Command,
) *Worker {
	worker := Worker{
		Container: c,
		Config:    config,
		Logger:    logger,
		CMD:       cmd,
		ENVPrefix: envPrefix,
	}
	worker.Bind()
	return &worker
}

func (w *Worker) Bind() {
	w.BindViper()
}

func (w *Worker) BindViper() {
	// bindFlags(w.CMD, w.Config)

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
}

func bindFlags(cmd *cobra.Command, cfg *viper.Viper) error {
	if err := cfg.BindPFlags(cmd.Flags()); err != nil {
		return err
	}
	for _, subCmd := range cmd.Commands() {
		if err := bindFlags(subCmd, cfg); err != nil {
			return err
		}
	}
	return nil
}
