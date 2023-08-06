package htmx_test

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/larschri/go-htmx"
)

// exampleHandler writes a greating and triggers myevent for an htmx request.
func exampleHandler(w http.ResponseWriter, r *http.Request) {
	if !htmx.HXRequest.Is(r) {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}
	htmx.HXTrigger.Set(w, "myevent")
	fmt.Fprintf(w, "Hello %s", htmx.HXCurrentURL.Get(r))
}

func TestExampleHandler(t *testing.T) {
	req := httptest.NewRequest("GET", "http://example.com/greating", nil)
	req.Header.Add("HX-Request", "true")
	req.Header.Add("HX-Current-URL", "http://example.com/index.html")

	recorder := httptest.NewRecorder()
	exampleHandler(recorder, req)
	response := recorder.Result()

	event := response.Header.Get("HX-Trigger")
	if event != "myevent" {
		t.Errorf("expected myevent, got %s", event)
	}

	bs, err := io.ReadAll(response.Body)
	if string(bs) != "Hello http://example.com/index.html" {
		t.Errorf("got %#v (err:%v)", string(bs), err)
	}
}
