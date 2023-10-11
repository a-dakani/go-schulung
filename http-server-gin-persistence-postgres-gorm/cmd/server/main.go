package main

import (
	"fmt"
	"github.com/a-dakani/go-schulung/http-server-gin-persistence-postgres-gorm/http"
	"github.com/a-dakani/go-schulung/http-server-gin-persistence-postgres-gorm/postgres"
	"golang.org/x/net/context"
)

func main() {
	repository, _ := postgres.NewAutoRepository(context.Background())
	fmt.Println(http.StartServer(repository))
}
