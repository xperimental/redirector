package main // import "github.com/xperimental/redirector"

import (
	"flag"
	"log"
	"net/http"
)

func main() {
	if err := checkArgs(); err != nil {
		log.Println(err)
		flag.Usage()
		return
	}

	r, err := makeRedirector(destination)
	if err != nil {
		log.Printf("Error in destination format: %s", err)
		return
	}

	log.Printf("Redirecting to %s", destination)
	http.Handle("/", r)

	log.Printf("Listen on %s", addr)
	log.Fatal(http.ListenAndServe(addr, nil))
}
