package main

import (
	"fmt"
	"github.com/a-dakani/go-schulung/contextAndLogging/ginserver"
	"github.com/a-dakani/go-schulung/contextAndLogging/http"
	autoRepository "github.com/a-dakani/go-schulung/contextAndLogging/memory"
)

func main() {
	repository := &autoRepository.AutoRepository{
		Autos: make([]ginserver.Auto, 0),
	}
	fmt.Println(http.StartServer(repository))
}
