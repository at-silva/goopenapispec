package goopenapispec

type MediaType struct {
	// Schema defining the type used for the request or response body.
	Schema *Schema `json:"schema,omitempty" yaml:"schema,omitempty"`

	// Encoding A map of possible representations for the media type.
	// The key is a media type and the value describes it.
	// For example, "application/json": {...}.
	Encoding map[string]*Encoding `json:"encoding,omitempty" yaml:"encoding,omitempty"`
}
