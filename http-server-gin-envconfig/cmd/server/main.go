package main

import (
	"fmt"
	"github.com/a-dakani/go-schulung/http-server-gin-envconfig/http"
	"github.com/a-dakani/go-schulung/http-server-gin-envconfig/postgres"
	"github.com/kelseyhightower/envconfig"
	"golang.org/x/net/context"
)

func main() {
	applConfig := http.Config{}
	err := envconfig.Process("", applConfig)
	if err != nil {
		panic(err)
	}
	repository, _ := postgres.NewAutoRepository(context.Background())
	fmt.Println(http.StartServer(applConfig, repository))
}
