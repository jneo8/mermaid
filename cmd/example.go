package main

import (
	"github.com/gin-gonic/gin"
	"github.com/jneo8/mermaid"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"go.uber.org/dig"
	"os"
)

type Service interface {
	Run() error
}

type HttpService struct {
	Engine   *gin.Engine
	Username string
}

func (s *HttpService) Run() error {
	s.Engine.GET("ping", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "pong"})
	})

	s.Engine.GET("user", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": s.Username})
	})
	return s.Engine.Run()
}

type ServiceOptions struct {
	dig.In
	Username string `name:"username"`
}

func NewHttpService(opts ServiceOptions) Service {
	log.Infof("New HpptService %v", opts)
	return &HttpService{
		Engine:   gin.Default(),
		Username: opts.Username,
	}
}

var rootCMD = &cobra.Command{
	Use:   "exmaple",
	Short: "exmaple",
	RunE: func(cmd *cobra.Command, args []string) error {

		initializers := []interface{}{
			NewHttpService,
		}
		runable := func(service Service, logger *log.Logger, config *viper.Viper) error {
			logger.Infof("Config: %v", config.AllSettings())
			if err := service.Run(); err != nil {
				logger.Error(err)
			}
			return nil
		}
		return mermaid.Run(cmd, runable, initializers...)
	},
}

func init() {
	rootCMD.Flags().String("username", "userA", "username")
}

func main() {
	if err := rootCMD.Execute(); err != nil {
		log.Error(err)
		os.Exit(1)
	}
}
