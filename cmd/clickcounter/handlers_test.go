package main

import (
	"github.com/humanitec/webinar-click-counter/internal/campaign"
	"github.com/steinfletcher/apitest"
	"net/http"
	"testing"
)

func TestRedirect(t *testing.T) {
	conf := campaign.NewSettings()
	svc := campaign.NewClickService(conf)

	apitest.New().
		HandlerFunc(redirect(conf, svc)).
		Get("/").
		Query("o", "https://geraldoandrade.com").
		Query("d", "https://humanitec.com").
		Expect(t).
		Status(http.StatusFound).
		End()
}
