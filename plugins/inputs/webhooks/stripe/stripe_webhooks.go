package stripe

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/influxdata/telegraf"
)

type StripeWebhook struct {
	Path string
	acc  telegraf.Accumulator
}

func (st *StripeWebhook) Register(router *mux.Router, acc telegraf.Accumulator) {
	router.HandleFunc(st.Path, st.eventHandler).Methods("POST")
	log.Printf("Started the webhooks_stripe on %s\n", st.Path)
	st.acc = acc
}

func (st *StripeWebhook) eventHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusBadRequest)
}
