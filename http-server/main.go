package main

import (
	"log"
	"net/http"
)

func main() {
	// create a simple server with one handler

	log.Println("Starting server on port 8080")
	targetHandlerWithOnlyGet := http.HandlerFunc(handler)
	http.Handle("/only-get", onlyGet(targetHandlerWithOnlyGet))
	http.Handle("/auth", onlyGet(basicAuth(targetHandlerWithOnlyGet)))
	log.Fatal(http.ListenAndServe("localhost:8080", nil))

}

func onlyGet(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "GET" {
			log.Println("Method not allowed")
			w.WriteHeader(http.StatusMethodNotAllowed)
		} else {
			next.ServeHTTP(w, r)
		}
	})
}

func basicAuth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// get the basic auth credentials
		user, pass, ok := r.BasicAuth()

		if !ok {
			w.Header().Add("WWW-Authenticate", `Basic realm="GO"`)
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		log.Printf("User: %s Pass: %s", user, pass)
	})
}

func handler(w http.ResponseWriter, r *http.Request) {
	log.Println("Received request from", r.RemoteAddr)
	w.Write([]byte("Hello World!"))
}
