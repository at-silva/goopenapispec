package goopenapispec

type License struct {
	// Name represents the name of the license type.
	Name string `json:"name" yaml:"name"`

	// URL contains the URL pointing to the license terms.
	URL string `json:"url,omitempty" yaml:"url,omitempty"`
}
