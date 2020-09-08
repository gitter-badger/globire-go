package api

import (
	"context"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"

	"github.com/pkg/errors"
)

const defaultURL = "https://api.companieshouse.gov.uk"

// API is provides all functionality of the Companies House REST API
type API struct {
	Key string
	URL *url.URL
}

// New returns an initialized instance of an API.
func New(apiKey string) (*API, error) {
	if apiKey == "" {
		return nil, fmt.Errorf("empty API key")
	}

	u, err := url.Parse(defaultURL)
	if err != nil {
		return nil, errors.Wrap(err, "parsing URL")
	}

	return &API{Key: apiKey, URL: u}, nil
}

// RequestError represents an error returned by the CH API.
type RequestError struct {
	Errors []struct {
		Error        string            `json:"error"`
		ErrorValues  map[string]string `json:"error_values"`
		Location     string            `json:"location"`
		LocationType string            `json:"location_type"`
		Type         string            `json:"type"`
	} `json:"errors"`
}

// RequestError implements the Error interface and returns the first error.
func (err *RequestError) Error() string {
	return err.Errors[0].Error
}

// IsRequestError returns true if the provided error is of type RequestError, as well as the asserted error.
func IsRequestError(err error) (bool, *RequestError) {
	e, ok := err.(*RequestError)
	return ok, e
}

// DoRequest makes a request to the API and returns the raw http resonse.
// if the API doesn't return statusOK, DoRequest will try to decode the error to a RequestError, or
// return the status description if decoding fails.
func (a *API) DoRequest(ctx context.Context, method string, path string, params url.Values, body io.Reader) (*http.Response, error) {
	u, err := url.Parse(a.URL.String() + path)
	if err != nil {
		return nil, errors.Wrap(err, "parsing URL")
	}

	if params != nil {
		u.RawQuery = params.Encode()
	}

	req, err := http.NewRequestWithContext(ctx, method, u.String(), body)
	if err != nil {
		return nil, errors.Wrap(err, "creating reuqest")
	}

	if a.Key == "" {
		return nil, fmt.Errorf("empty API key")
	}
	req.SetBasicAuth(a.Key, "")

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := http.Client{Transport: tr}

	resp, err := client.Do(req)
	if err != nil {
		return nil, errors.Wrap(err, "http request")
	}

	if resp.StatusCode != http.StatusOK {
		defer resp.Body.Close()
		var err RequestError

		if err := json.NewDecoder(resp.Body).Decode(&err); err != nil {
			return nil, fmt.Errorf("%v", resp.Status)
		}
	}

	return resp, nil
}

// Do makes a request to the API and decoded the result into v.
// v should be a pointer.
func (a *API) Do(ctx context.Context, method string, path string, params url.Values, body io.Reader, v interface{}) error {
	resp, err := a.DoRequest(ctx, method, path, params, body)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if err := json.NewDecoder(resp.Body).Decode(v); err != nil {
		return errors.Wrap(err, "decoding response")
	}

	return nil
}
