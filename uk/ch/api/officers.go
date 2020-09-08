package api

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"strings"

	"github.com/appinesshq/globire-go/uk/ch/api/enum"
)

// IdentificationType represents identification type
type IdentificationType string

// String implements the Stringer interface to get a human readable string from the CH enums
func (f IdentificationType) String() string {
	return enum.Constants.Get("identification_type", string(f))
}

// OfficerRole represents an officer's role
type OfficerRole string

// String implements the Stringer interface to get a human readable string from the CH enums
func (f OfficerRole) String() string {
	return enum.Constants.Get("officer_role", string(f))
}

type (
	// OfficerDateOfBirth struct consists of Day(int), Month (int) and Year (int)
	OfficerDateOfBirth struct {
		Day   int `json:"day"`
		Month int `json:"month"`
		Year  int `json:"year"`
	}

	// Identification represents the details of a form of identification
	Identification struct {
		IdentificationType IdentificationType `json:"identification_type"`
		Authority          string             `json:"legal_authority"`
		LegalForm          string             `json:"legal_form"`
		PlaceRegistered    string             `json:"place_registered"`
		RegistrationNumber string             `json:"registration_number"`
	}

	// Officers contains the server response of a data request to the companies house API
	Officers struct {
		Etag          string    `json:"etag"`
		Kind          string    `json:"kind"`
		Start         int       `json:"start_index"`
		ItemsPerPage  int       `json:"items_per_page"`
		TotalResults  int       `json:"total_results"`
		ActiveCount   int       `json:"active_count"`
		InactiveCount int       `json:"inactive_count"`
		ResignedCount int       `json:"resigned_count"`
		Items         []Officer `json:"items"`
		Links         struct {
			Self string `json:"self"`
		} `json:"Links"`
	}
)

// Officer struct contains the data of a company's officers
type Officer struct {
	Address            Address            `json:"address"`
	AppointedOn        ChDate             `json:"appointed_on"`
	CountryOfResidence string             `json:"country_of_residence"`
	DateOfBirth        OfficerDateOfBirth `json:"date_of_birth"`
	FormerNames        []struct {
		Forenames string `json:"forenames"`
		Surname   string `json:"surname"`
	} `json:"former_names"`
	Identification Identification `json:"identification"`
	Links          struct {
		Officer struct {
			Appointments string `json:"appointments"`
		} `json:"officer"`
	} `json:"links"`
	Name        string      `json:"name"`
	Nationality string      `json:"nationality"`
	Occupation  string      `json:"occupation"`
	OfficerRole OfficerRole `json:"officer_role"`
	ResignedOn  ChDate      `json:"resigned_on"`
}

// ID returns an officer's ID
func (o Officer) ID() string {
	a := strings.Split(o.Links.Officer.Appointments, "/")
	return a[2]
}

// Officers gets and return a company's officers
// Possible options: ItemsPerPage. OfficerType, RegisterView, StartIndex, OrderBy
func (c *Company) Officers(options ...Option) (*Officers, error) {
	// Prepare the response
	res := Officers{}
	params := url.Values{}
	for _, option := range options {
		option(&params)
	}

	// Make a call to the service
	path := fmt.Sprintf("/company/%s/officers", c.CompanyNumber)
	err := c.api.Do(context.Background(), http.MethodGet, path, params, nil, &res)
	// Ensure to close the ReadCloser
	// defer b.Close()
	if err != nil {
		return nil, err
	}

	// Decode the service's response to a Company struct and return the result or error
	// if err := json.NewDecoder(b).Decode(&res); err != nil {
	// 	return nil, err
	// }
	return &res, nil
}
