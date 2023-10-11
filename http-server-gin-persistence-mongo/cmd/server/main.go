package main

import (
	"fmt"
	"github.com/a-dakani/go-schulung/http-server-gin-persistence-mongo/http"
	"github.com/a-dakani/go-schulung/http-server-gin-persistence-mongo/mongodb"
	"golang.org/x/net/context"
)

func main() {

	repository, _ := mongodb.NewAutoRepository(context.Background())
	fmt.Println(http.StartServer(repository))
}
