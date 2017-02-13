package stripe

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"io/ioutil"
	"time"

	"github.com/gorilla/mux"
	"github.com/influxdata/telegraf"
)

type StripeWebhook struct {
	Path string
	acc  telegraf.Accumulator
}

func (st *StripeWebhook) Register(router *mux.Router, acc telegraf.Accumulator) {
	router.HandleFunc(st.Path, st.eventHandler).Methods("POST")
	log.Printf("I! Started the webhooks_stripe on %s\n", st.Path)
	st.acc = acc
}

func (st *StripeWebhook) eventHandler(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	event := &Event{}
	err = json.Unmarshal(body, event)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	object, err := NewEvent(event)
	if err != nil {
		w.WriteHeader(http.StatusOK)
		return
	}

	tags := map[string]string{
		"event": event.Type,
	}

	st.acc.AddFields("stripe_webhooks", object.Fields(), tags, time.Now())

	w.WriteHeader(http.StatusOK)
}

func generateObject(object Object, data []byte) (Object, error) {
	err := json.Unmarshal(data, object)
	if err != nil {
		return nil, err
	}
	return object, nil
}

func NewEvent(event *Event) (Object, error) {
	switch event.Type {
	case "charge.succeeded":
		return generateObject(&Charge{}, event.Data.Object)
	case "charge.failed":
		return generateObject(&Charge{}, event.Data.Object)
	case "charge.refunded":
		return generateObject(&Charge{}, event.Data.Object)
	case "customer.subscription.created":
		return generateObject(&Subscription{}, event.Data.Object)
	case "customer.subscription.updated":
		return generateObject(&Subscription{}, event.Data.Object)
	case "customer.subscription.deleted":
		return generateObject(&Subscription{}, event.Data.Object)
	default:
		return nil, errors.New("Not implemented type: " + event.Type)
	}
}
