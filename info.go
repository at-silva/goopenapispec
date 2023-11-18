package goopenapispec

type Info struct {
	// Title represents the title of the API.
	Title string `json:"title" yaml:"title"`

	// Description provides a brief description of the API.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`

	// TermsOfService contains the URL or information regarding the terms of service.
	TermsOfService string `json:"termsOfService,omitempty" yaml:"termsOfService,omitempty"`

	// Contact holds information about the contact person or organization for the API.
	Contact Contact `json:"contact,omitempty" yaml:"contact,omitempty"`

	// License contains information about the API's license.
	License License `json:"license,omitempty" yaml:"license,omitempty"`

	// Version specifies the version of the API.
	Version string `json:"version" yaml:"version"`
}
