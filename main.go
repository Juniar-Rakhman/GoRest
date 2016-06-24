package main

import (
	"log"
	"net/http"
	"jrakhman/router"
)

func main() {
	router := router.NewRouter()
	log.Fatal(http.ListenAndServe(":3000", router))
}