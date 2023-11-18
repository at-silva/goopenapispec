package goopenapispec

type ResponseRef struct {
	Response

	// Ref is a reference to another response definition.
	Ref string `json:"$ref,omitempty" yaml:"$ref,omitempty"`
}
