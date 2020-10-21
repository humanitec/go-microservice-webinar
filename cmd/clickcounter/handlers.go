package main

import (
	"fmt"
	"github.com/humanitec/webinar-click-counter/internal/campaign"
	"log"
	"net/http"
	"time"
)

func redirect(_ *campaign.Settings, svc *campaign.ClickService) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		// get the data from the request
		origin := r.URL.Query().Get("o")
		destination := r.URL.Query().Get("d")

		data := campaign.Click{
			Origin:      origin,
			Destination: destination,
			Timestamp:   time.Now().UTC(),
		}

		// save the click data
		//go func() {
		if err := svc.AddClick(data); err != nil {
			log.Printf("error storing click information: %v", fmt.Errorf("%w", err))
			w.WriteHeader(500)
			return
		}
		//}()

		// response with the redirect
		w.Header().Add("Location", destination)
		w.WriteHeader(http.StatusFound)
		log.Printf("redirecting from %s to %s", origin, destination)
	}
}

// live Kubernetes automated liveness checks
func live(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)
}

// ready Kubernetes automated readiness checks
func ready(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)
}
