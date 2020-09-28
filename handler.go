package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/RedeployAB/gofunc/functions"
)

func main() {
	port := os.Getenv("FUNCTIONS_CUSTOMHANDLER_PORT")
	if len(port) == 0 {
		port = "8080"
	}

	r := http.NewServeMux()
	r.HandleFunc("/IncomingHTTP", functions.IncomingHTTP)
	fmt.Println("Listening on port:", port)
	log.Fatal(http.ListenAndServe(":"+port, r))
}
