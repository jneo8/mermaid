# Mermaid

![testing](https://github.com/jneo8/mermaid/workflows/testing/badge.svg)

Mermaid is a tool helping user use dependency injection when using [Cobra](https://github.com/spf13/cobra).

## What Mermaid do?

Mermaid bind flags from cobra to viper as global settings. And provide all settings to container automatically. 


## Example

```go
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

var cmd = &cobra.Command{
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
	cmd.Flags().String("username", "userA", "username")
}

func main() {
	if err := cmd.Execute(); err != nil {
		log.Error(err)
		os.Exit(1)
	}
}
```

## Thanks

The basic idea isn't from me. Is from here [結構化的 Go 專案設計模式與風格 — 4 - getamis - Medium](https://medium.com/getamis/%E7%B5%90%E6%A7%8B%E5%8C%96%E7%9A%84-go-%E5%B0%88%E6%A1%88%E8%A8%AD%E8%A8%88%E6%A8%A1%E5%BC%8F%E8%88%87%E9%A2%A8%E6%A0%BC-2-548fec8cd9bb).
