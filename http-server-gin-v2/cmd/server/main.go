package main

import (
	"fmt"
	"github.com/a-dakani/go-schulung/http-server-gin/v2/ginserver"
	"github.com/a-dakani/go-schulung/http-server-gin/v2/http"
	autoRepository "github.com/a-dakani/go-schulung/http-server-gin/v2/memory"
)

func main() {
	repository := &autoRepository.AutoRepository{
		Autos: make([]ginserver.Auto, 0),
	}
	fmt.Println(http.StartServer(repository))
}
