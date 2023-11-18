package goopenapispec

type ResponseHeader struct {
	// Description provides additional information about the response header.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`

	// Required specifies whether the response header is mandatory.
	Required bool `json:"required,omitempty" yaml:"required,omitempty"`

	// Deprecated indicates if the response header is deprecated.
	Deprecated bool `json:"deprecated,omitempty" yaml:"deprecated,omitempty"`

	// Schema contains the schema defining the structure of the response header.
	Schema *Schema `json:"schema,omitempty" yaml:"schema,omitempty"`

	// Example holds an example value for the response header.
	Example interface{} `json:"example,omitempty" yaml:"example,omitempty"`

	// Examples contain multiple named examples for the response header.
	Examples map[string]*Example `json:"examples,omitempty" yaml:"examples,omitempty"`
}
