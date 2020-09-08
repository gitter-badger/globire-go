package tests

import (
	"net/http"
	"net/http/httptest"
	"strings"
)

const (
	companyData = `{
		"links": {
		  "filing_history": "/company/12345678/filing-history",
		  "self": "/company/12345678",
		  "officers": "/company/12345678/officers",
		  "persons_with_significant_control": "/company/12345678/persons-with-significant-control"
		},
		"has_charges": false,
		"date_of_creation": "2019-06-25",
		"accounts": {
		  "next_accounts": {
			"period_start_on": "2019-06-25",
			"due_on": "2021-06-25",
			"overdue": false,
			"period_end_on": "2020-06-30"
		  },
		  "next_made_up_to": "2020-06-30",
		  "overdue": false,
		  "accounting_reference_date": {
			"month": "06",
			"day": "30"
		  },
		  "last_accounts": {
			"type": "null"
		  },
		  "next_due": "2021-06-25"
		},
		"company_number": "12345678",
		"company_status": "active",
		"company_name": "TEST LTD",
		"confirmation_statement": {
		  "next_made_up_to": "2020-06-24",
		  "overdue": false,
		  "next_due": "2020-08-05"
		},
		"registered_office_address": {
		  "address_line_2": "15 Test Road",
		  "locality": "Test Town",
		  "postal_code": "TS1 2TS",
		  "country": "United Kingdom",
		  "address_line_1": "Office 1"
		},
		"registered_office_is_in_dispute": false,
		"sic_codes": [
		  "58290",
		  "62012"
		],
		"etag": "b400b09dd02caf1c3a54ea40b8672637f664bf49",
		"has_insolvency_history": false,
		"undeliverable_registered_office_address": false,
		"jurisdiction": "england-wales",
		"type": "ltd",
		"can_file": true
	  }`

	officerData = `{
		"total_results": 1,
		"resigned_count": 0,
		"start_index": 0,
		"inactive_count": 0,
		"etag": "43ae73fcee9a9c92f19727bbd4feb8daead859ba",
		"kind": "officer-list",
		"active_count": 1,
		"items": [
		  {
			"officer_role": "director",
			"date_of_birth": {
			  "year": 1977,
			  "month": 12
			},
			"nationality": "Dutch",
			"name": "PERSON, Test",
			"appointed_on": "2019-06-25",
			"country_of_residence": "Lithuania",
			"occupation": "Company Director",
			"address": {
			  "premises": "1",
			  "postal_code": "TS1 T1N",
			  "locality": "Test Town",
			  "country": "United Kingdom",
			  "address_line_1": "Test Road",
			  "address_line_2": "Office 1"
			},
			"links": {
			  "officer": {
				"appointments": "/officers/e4-ScyHpxNNUh6ZyV9wnqZS1kfY/appointments"
			  }
			}
		  }
		],
		"links": {
		  "self": "/company/12345678/officers"
		},
		"items_per_page": 35
	  }`
)

// NewMockServer simulates the API for testing purposes.
// Supported requests:
// 12345678 - Active Limited company
// Other company numbers - Not found error
func NewMockServer() *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		path := strings.Split(r.URL.Path[1:], "/")
		switch path[0] {
		case "company":
			getCompany(w, path)
		default:
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("Invalid request path"))
		}
	}))
}

func getCompany(w http.ResponseWriter, path []string) {
	switch {
	case len(path) >= 2 && path[1] == "12345678":
		switch {
		case len(path) > 2 && path[2] == "officers":
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(officerData))
		default:
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(companyData))
		}

	default:
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Not found"))
	}
}
