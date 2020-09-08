package api

import (
	"fmt"
	"net/url"
	"strconv"
	"strings"
)

// Option is a type used for passing options to modify an API call's parameters
type Option func(*url.Values)

// ItemsPerPage limits the amount of pages that will be returned
func ItemsPerPage(val int) Option {
	return func(v *url.Values) {
		v.Set("items_per_page", strconv.Itoa(val))
	}
}

// StartIndex sets the result's offset
func StartIndex(val int) Option {
	return func(v *url.Values) {
		v.Set("start_index", strconv.Itoa(val))
	}
}

// officerType represents a type of officers which can be searched on
type officerType string

const (
	// Directors are company directors
	Directors officerType = "directors"

	// Secretaries are company secretaries
	Secretaries officerType = "secretaries"

	// LLPMembers are all kinds of LLP members
	LLPMembers officerType = "llp-members"
)

// OfficerType allows to select the type of officer to search for
// Using this function will automatically set register_view to true as well, as this is required by the API
func OfficerType(val officerType) Option {
	return func(v *url.Values) {
		v.Set("register_type", string(val))
		v.Set("register_view", strconv.FormatBool(true))
	}
}

// RegisterView limits the result to contain only officers with a full date of birth
func RegisterView(val bool) Option {
	return func(v *url.Values) {
		v.Set("register_view", strconv.FormatBool(val))
	}
}

// orderBy are the options which can be used for sorting the results
type orderBy string

const (
	// AppointedOn is used for sorting on the appointment date
	AppointedOn orderBy = "appointed_on"

	// ResignedOn is used for sorting on the date of resignation
	ResignedOn orderBy = "resigned_on"

	// Surname is used for sorting on an officer's surname
	Surname orderBy = "surname"
)

// OrderBy is an option used for sorting the results
// Takes in an orderBy value and a boolean value for sorting descending/ascending
func OrderBy(val orderBy, desc bool) Option {
	return func(v *url.Values) {
		if desc {
			v.Set("order_by", fmt.Sprintf("-%s", val))
		} else {
			v.Set("order_by", string(val))
		}
	}
}

// Category are categories to filter by (inclusive)
func Category(val ...string) Option {
	return func(v *url.Values) {
		for i := range val {
			val[i] = strings.ToLower(val[i])
		}
		v.Set("category", strings.Join(val, ","))
	}
}
