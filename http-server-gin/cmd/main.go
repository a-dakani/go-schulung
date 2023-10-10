package main

import (
	"fmt"
	"github.com/a-dakani/go-schulung/http-server-gin/http"
	"github.com/a-dakani/go-schulung/http-server-gin/memory"
)

func main() {
	repository := &autoRepository.AutoRepository{}
	fmt.Println(http.StartServer(repository))
}
