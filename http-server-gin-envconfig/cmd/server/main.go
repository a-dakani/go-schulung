package main

import (
	"fmt"
	"github.com/a-dakani/go-schulung/http-server-gin-envconfig/http"
	"github.com/a-dakani/go-schulung/http-server-gin-envconfig/postgres"
	"github.com/joho/godotenv"
	"golang.org/x/net/context"
	"os"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}
	repository, _ := postgres.NewAutoRepository(context.Background())
	fmt.Println(http.StartServer(fmt.Sprintf("%s:%s", os.Getenv("HOST"), os.Getenv("PORT")), repository))
}
