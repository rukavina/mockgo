package main

import (
	"flag"
	"log"
	"net/http"
	"strconv"
)

func main() {
	portPtr := flag.Int("port", 8082, "port number")
	flag.Parse()

	r, err := load("config.json")
	if err != nil {
		log.Fatalf("Error loading config file: %s", err)
	}

	port := strconv.Itoa(*portPtr)
	log.Printf("Running on port [%s]\n", port)
	log.Fatal(http.ListenAndServe(":"+port, r))
}
