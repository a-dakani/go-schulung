package main

import (
	"context"
	"encoding/json"
	"github.com/a-dakani/go-schulung/http-server-gin-kafka/ginserver"
	"io"
	"log"
	"net/http"
)

func main() {
	ctx := context.Background()
	request, err := http.NewRequestWithContext(ctx, http.MethodGet, "http://localhost:8080/api/autos", nil)
	if err != nil {
		return
	}
	request.SetBasicAuth("foo", "bar")
	//request.Header.Set("Accept", "application/xml")
	request.Header.Set("Accept", "application/json")
	request.Header.Set("x-trace-id", "1234")

	response, err := http.DefaultClient.Do(request)
	if err != nil {
		log.Fatal(err)
	}
	if response.StatusCode != http.StatusOK {
		log.Fatal(http.StatusText(response.StatusCode))
	} else {
		log.Println("Status code is 200")

	}

	defer response.Body.Close()
	all, err := io.ReadAll(response.Body)
	if err != nil {
		return
	}

	var autos []ginserver.Audi
	err = json.Unmarshal(all, &autos)
	if err != nil {
		return
	}

	log.Println(len(autos))
	for _, auto := range autos {
		log.Printf("%+v", auto)
		log.Printf("%+v", auto.Getriebe)
	}

}
