//************************************************************************//
// API "cellar": Application Media Types
//
// Generated with goagen v1.0.0, command line:
// $ goagen
// --design=github.com/ivan3bx/medal-service/design
// --out=$(GOPATH)/src/github.com/ivan3bx/medal-service
// --version=v1.0.0
//
// The content of this file is auto-generated, DO NOT MODIFY
//************************************************************************//

package app

import "github.com/goadesign/goa"

// A list of country medals (default view)
//
// Identifier: application/vnd.medals+json
type Medals struct {
	// Number of bronze medals
	Bronze int `form:"bronze" json:"bronze" xml:"bronze"`
	// Short identifier for country
	CountryID string `form:"countryId" json:"countryId" xml:"countryId"`
	// Descriptive name for country
	CountryName *string `form:"countryName,omitempty" json:"countryName,omitempty" xml:"countryName,omitempty"`
	// Number of gold medals
	Gold int `form:"gold" json:"gold" xml:"gold"`
	// Number of silver medals
	Silver int `form:"silver" json:"silver" xml:"silver"`
}

// Validate validates the Medals media type instance.
func (mt *Medals) Validate() (err error) {
	if mt.CountryID == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "countryId"))
	}

	return
}

// MedalsCollection is the media type for an array of Medals (default view)
//
// Identifier: application/vnd.medals+json; type=collection
type MedalsCollection []*Medals

// Validate validates the MedalsCollection media type instance.
func (mt MedalsCollection) Validate() (err error) {
	for _, e := range mt {
		if e.CountryID == "" {
			err = goa.MergeErrors(err, goa.MissingAttributeError(`response[*]`, "countryId"))
		}

	}
	return
}
