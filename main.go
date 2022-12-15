package main

import (
	"log"
	"net/http"

	"github.com/hendralatumeten/cakestore/src/routers"
)

func main() {
	con, err := routers.New()
	if err != nil {
		panic(err)
	}

	log.Println("Listening on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", con))

}
