package main

import (
	"log"
	"net/http"
)

func redirect(w http.ResponseWriter, r *http.Request) {
	// get the data from the request
	origin := r.URL.Query().Get("o")
	destination := r.URL.Query().Get("d")

	// save the data


	// response with the redirect
	w.Header().Add("Location", "https://geraldoandrade.com")
	w.WriteHeader(302)
	log.Print("Redirecting")
}

func favicon(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)
}

// live Kubernetes automated liveness checks
func live(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)
}

// ready Kubernetes automated readiness checks
func ready(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)
}
