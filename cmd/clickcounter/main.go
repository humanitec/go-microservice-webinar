package main

import (
	"fmt"
	"github.com/humanitec/webinar-click-counter/internal/campaign"
	"log"
	"net/http"
)

func main() {
	conf := campaign.NewSettings()
	svc := campaign.NewClickService(conf)

	if err := svc.Init(); err != nil {
		log.Fatalf("could not initialize the elasticsearch mappings: %v", fmt.Errorf("%w", err))
	}

	http.HandleFunc("/", redirect(conf, svc))
	http.HandleFunc("/ready", ready)
	http.HandleFunc("/live", live)

	log.Printf("app initialized at port %v", conf.Port)
	if err := http.ListenAndServe(":9000", nil); err != nil {
		log.Fatalf("server error: %v", fmt.Errorf("%w", err))
	}
}
