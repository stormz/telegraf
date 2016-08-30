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

func TestBadRequest(t *testing.T) {
	var acc testutil.Accumulator
	st := &StripeWebhook{Path: "/stripe", acc: &acc}
	resp := postWebhooks(st, "")
	if resp.Code != http.StatusBadRequest {
		t.Errorf("POST returned HTTP status code %v.\nExpected %v", resp.Code, http.StatusOK)
	}
}
