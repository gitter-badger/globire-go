package api_test

import (
	"context"
	"net/http"
	"os"
	"testing"

	ch "github.com/appinesshq/globire-go/uk/ch/api"
)

func TestNewAPI(t *testing.T) {
	api, err := ch.New("test")
	if err != nil {
		t.Fatalf("expected to pass, but got %v", err)
	}

	if got, expected := api.Key, "test"; got != expected {
		t.Errorf("expected key to be %q, but got %q", expected, got)
	}

	if got, expected := api.URL.String(), "https://api.companieshouse.gov.uk"; got != expected {
		t.Errorf("expected URL to be %q, but got %q", expected, got)
	}
}

func TestDoRequest(t *testing.T) {
	key := os.Getenv("CH_API_TEST_KEY")
	if key == "" {
		t.Log("skipping TestDoRequest, because CH_API_TEST_KEY isn't set")
		return
	}

	api, err := ch.New(key)
	if err != nil {
		t.Fatalf("expected to pass, but got %v", err)
	}

	_, err = api.DoRequest(context.Background(), http.MethodGet, "/company/12068026", nil, nil)
	if err != nil {
		t.Fatalf("expected to pass, but got %v", err)
	}
}
