package main

import (
	"log"
	"net/http"

	"github.com/zuils/pokepaste"
)

func main() {
	log.Fatal(http.ListenAndServe(":8080", pokepaste.Handler))
}
