package main

import (
	"fmt"
	"net/http"
)

func main() {
	server := &http.Server{
		Addr:    "3000",
		Handler: http.HandlerFunc(basicHandler),
	}
	err := server.ListenAndServe()
	if err != nil {
		fmt.Println("Fuck u mate")
	}
}

func basicHandler(wr http.ResponseWriter, r *http.Request) {

}
