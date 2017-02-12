package stripe

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/influxdata/telegraf/testutil"
)

func postWebhooks(st *StripeWebhook, eventBody string) *httptest.ResponseRecorder {
	req, _ := http.NewRequest("POST", "/", strings.NewReader(eventBody))
	w := httptest.NewRecorder()
	w.Code = 500

	st.eventHandler(w, req)

	return w
}

func request(t *testing.T, eventBody string, expectedStatus int) testutil.Accumulator {
	var acc testutil.Accumulator
	st := &StripeWebhook{Path: "/stripe", acc: &acc}
	resp := postWebhooks(st, eventBody)
	if resp.Code != expectedStatus {
		t.Errorf("POST returned HTTP status code %v.\nExpected %v", resp.Code, expectedStatus)
	}
	return acc
}


func TestBadRequest(t *testing.T) {
	request(t, "", http.StatusBadRequest)
}

func TestChargeSucceeded(t *testing.T) {
	var acc = request(t, EventChargeJSON("charge.succeeded"), http.StatusOK)

	fields := map[string]interface{}{
		"id": "ch_19lEGwI4amPefGV6kCxtZyQC",
		"amount": 1800,
		"currency": "eur",
	}

	tags := map[string]string{
		"event": "charge.succeeded",
	}

	acc.AssertContainsTaggedFields(t, "stripe_webhooks", fields, tags)
}

func TestChargeFailed(t *testing.T) {
	var acc = request(t, EventChargeJSON("charge.failed"), http.StatusOK)

	fields := map[string]interface{}{
		"id": "ch_19lEGwI4amPefGV6kCxtZyQC",
		"amount": 1800,
		"currency": "eur",
	}

	tags := map[string]string{
		"event": "charge.failed",
	}

	acc.AssertContainsTaggedFields(t, "stripe_webhooks", fields, tags)
}

func TestChargeRefunded(t *testing.T) {
	var acc = request(t, EventChargeJSON("charge.refunded"), http.StatusOK)

	fields := map[string]interface{}{
		"id": "ch_19lEGwI4amPefGV6kCxtZyQC",
		"amount": 1800,
		"currency": "eur",
	}

	tags := map[string]string{
		"event": "charge.refunded",
	}

	acc.AssertContainsTaggedFields(t, "stripe_webhooks", fields, tags)
}

func TestPing(t *testing.T) {
	request(t, EventPingJSON(), http.StatusOK)
}
