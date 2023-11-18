package goopenapispec

type ExternalDoc struct {
	// Description field holds a description providing more details about the external documentation.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`

	// URL field contains the URL pointing to the external documentation.
	URL string `json:"url" yaml:"url"`
}
