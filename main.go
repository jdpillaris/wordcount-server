package main

import (
	"log"
	"net/http"

	controller "github.com/jdpillaris/server/controller"
)

func main() {
	http.HandleFunc("/", controller.Hello)
	http.HandleFunc("/upload", controller.Upload)
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
