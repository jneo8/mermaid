# Mermaid

![testing](https://github.com/jneo8/mermaid/workflows/testing/badge.svg)
[![codecov](https://codecov.io/gh/jneo8/mermaid/branch/master/graph/badge.svg)](https://codecov.io/gh/jneo8/mermaid)
[![Go Report Card](https://goreportcard.com/badge/github.com/jneo8/mermaid)](https://goreportcard.com/report/github.com/jneo8/mermaid)
[![GoDoc](https://godoc.org/github.com/jneo8/mermaid?status.svg)](https://godoc.org/github.com/jneo8/mermaid)
[![Release](https://img.shields.io/github/release/jneo8/mermaid.svg?style=plastic)](https://github.com/jneo8/mermaid/releases)
[![TODOs](https://badgen.net/https/api.tickgit.com/badgen/github.com/jneo8/mermaid)](https://www.tickgit.com/browse?repo=github.com/jneo8/mermaid)

Mermaid is a tool helping user use [Cobra](https://github.com/spf13/cobra), [Viper](https://github.com/spf13/viper) and [dig](https://github.com/uber-go/dig) together.

## What Mermaid do?

Mermaid bind flags from cobra to viper as settings. And provide all settings to dig container automatically.
Make it easy to setup and write testing.

## Basic args

- `config`: config file path or filename.yaml

- `config_name`: if `config` field is blank, mermaid will find config yaml file in `.`, `/`, `$HOME`, `./config`. Default name is `config`.

- `log_level`: Set logger's log level. Default is info level.

## Example


### Basic use

cmd/example/main.go

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

cmd/example/main_test.go

```go
package main

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"net/http"
	"sync"
	"testing"
	"time"
)

func TestRootCMD(t *testing.T) {
	wg := sync.WaitGroup{}

	wg.Add(1)
	go func() {
		defer wg.Done()
		// Overwrite config using args.
		rootCMD.SetArgs(
			[]string{"--username", "userB"},
		)
		if err := rootCMD.Execute(); err != nil {
			t.Error(err)
		}
	}()

	time.Sleep(1 * time.Second) // Waiting for gin engineer setup.
	resp, err := http.Get("http://127.0.0.1:8080/user")
	if err != nil {
		t.Error(err)
	}
	defer resp.Body.Close()

	var result map[string]interface{}
	json.NewDecoder(resp.Body).Decode(&result)

	assert.Equal(t, result["message"], "userB")
}
```

### Use customized logger

cmd/customized_mermaid/main.go

```go
package main

import (
	"github.com/jneo8/mermaid"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func main() {
	worker, err := mermaid.NewMermaid(
		&cobra.Command{},
		viper.New(),
		log.New(),
		"",
	)
	if err != nil {
		log.Fatal(err)
	}

	runable := func(logger *log.Logger) error {
		logger.Info("Customized mermaid worker")
		return nil
	}
	initializers := []interface{}{}
	if err := worker.Execute(runable, initializers...); err != nil {
		log.Fatal(err)
	}
}
```

## Thanks

The basic idea isn't from me. Is from here [結構化的 Go 專案設計模式與風格 — 4 - getamis - Medium](https://medium.com/getamis/%E7%B5%90%E6%A7%8B%E5%8C%96%E7%9A%84-go-%E5%B0%88%E6%A1%88%E8%A8%AD%E8%A8%88%E6%A8%A1%E5%BC%8F%E8%88%87%E9%A2%A8%E6%A0%BC-2-548fec8cd9bb).
