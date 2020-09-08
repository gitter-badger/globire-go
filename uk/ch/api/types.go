package api

import (
	"strconv"
	"strings"
	"time"
)

// Address struct contains the details of addresses
type Address struct {
	Etag         string `json:"etag"`
	Premises     string `json:"premises"`
	AddressLine1 string `json:"address_line_1"`
	AddressLine2 string `json:"address_line_2"`
	Locality     string `json:"locality"`
	Region       string `json:"region"`
	PostalCode   string `json:"postal_code"`
	Country      string `json:"country"`
	CareOf       string `json:"care_of"`
	PoBox        string `json:"po_box"`
}

// ChDate is type which supports unmarshalling from CH json response to a Go time type
type ChDate struct {
	time.Time
}

// UnmarshalJSON implements the unmarshalling functionality
func (cd *ChDate) UnmarshalJSON(b []byte) (err error) {
	s := strings.Trim(string(b), "\"")
	if len(s) == 0 {
		return
	}
	cd.Time, err = time.Parse("2006-01-02", s)
	return
}

// DateOfBirth is a type which supports unmarshalling from CH json response to a Go time type
type DateOfBirth struct {
	time.Time
}

// UnmarshalJSON implements the unmarshalling functionality
func (dob *DateOfBirth) UnmarshalJSON(b []byte) (err error) {
	s := strings.Trim(string(b), "\"")
	dob.Time, err = time.Parse("2006-01-02T15:04:00", s)
	return
}

type strint int

func (v *strint) UnmarshalJSON(b []byte) error {
	s := strings.Trim(string(b), "\"")
	i, err := strconv.Atoi(s)
	if err != nil {
		return err
	}
	*v = strint(i)
	return nil
}

func (v *strint) Int() int {
	return int(*v)
}
