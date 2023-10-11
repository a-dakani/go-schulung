package main

import (
	"fmt"
	"github.com/a-dakani/go-schulung/http-server-gin-kafka/ginserver"
	"github.com/a-dakani/go-schulung/http-server-gin-kafka/http"
	"github.com/a-dakani/go-schulung/http-server-gin-kafka/kafka"
	"github.com/a-dakani/go-schulung/http-server-gin-kafka/postgres"
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
	notifier := kafka.NewAutoNotifier()
	as := ginserver.NewAutoService(notifier, repository)
	fmt.Println(http.StartServer(fmt.Sprintf("%s:%s", os.Getenv("HOST"), os.Getenv("PORT")), as))
}
