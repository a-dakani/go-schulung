package main

import (
	"fmt"
	"github.com/a-dakani/go-schulung/http-server-gin-testing/ginserver"
	"github.com/a-dakani/go-schulung/http-server-gin-testing/http"
	autoRepository "github.com/a-dakani/go-schulung/http-server-gin-testing/memory"
)

func main() {
	repository := &autoRepository.AutoRepository{
		Autos: make([]ginserver.Auto, 0),
	}
	fmt.Println(http.StartServer(repository))
}
