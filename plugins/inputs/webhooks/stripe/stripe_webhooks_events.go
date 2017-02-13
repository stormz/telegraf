package stripe

import (
	"encoding/json"
)

type Object interface {
	Tags() map[string]string
	Fields() map[string]interface{}
}

type EventData struct {
	Object  json.RawMessage `json:"object"`
}

type Event struct {
	Id      string    `json:"id"`
	Type    string    `json:"type"`
	Created int       `json:"created"`
	Data    EventData `json:"data"`
}

type Charge struct {
	Id       string `json:"id"`
	Amount   int    `json:"amount"`
	Currency string `json:"currency"`
}

func (charge *Charge) Tags() map[string]string {
	return map[string]string{
		"event": "charge.succeeded",
	}
}

func (charge *Charge) Fields() map[string]interface{} {
	return map[string]interface{}{
		"id": charge.Id,
		"amount": charge.Amount,
		"currency": charge.Currency,
	}
}

type Plan struct {
	Id string `json:"id"`
}

type Subscription struct {
	Id   string `json:"id"`
	Plan Plan   `json:"plan"`
}

func (subscription *Subscription) Tags() map[string]string {
	return map[string]string{}
}

func (subscription *Subscription) Fields() map[string]interface{} {
	return map[string]interface{}{
		"id": subscription.Id,
		"plan": subscription.Plan.Id,
	}
}
