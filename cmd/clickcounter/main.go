package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", redirect)
	http.HandleFunc("/favicon.ico", favicon)
	http.HandleFunc("/ready", ready)
	http.HandleFunc("/live", live)



	if err := http.ListenAndServe(":9000", nil); err != nil {
		log.Fatalf("server error: %v", fmt.Errorf("%w", err))
	}
}
