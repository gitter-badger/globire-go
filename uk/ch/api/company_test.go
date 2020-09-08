package api_test

import (
	"net/url"
	"testing"

	ch "github.com/appinesshq/globire-go/uk/ch/api"
	"github.com/appinesshq/globire-go/uk/ch/api/tests"
)

func TestGetCompany(t *testing.T) {
	api, err := ch.New("12345")
	if err != nil {
		t.Fatalf("expected to pass, but got: %v", err)
	}

	ts := tests.NewMockServer()
	api.URL, err = url.Parse(ts.URL)
	if err != nil {
		t.Fatalf("expected to pass, but got: %v", err)
	}

	c, err := api.GetCompany("12345678")
	if err != nil {
		t.Fatalf("expected to pass, but got: %v", err)
	}

	if got, expected := c.Name, "TEST LTD"; got != expected {
		t.Fatalf("expected %q, but got %q", expected, got)
	}
}
