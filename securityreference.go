package goopenapispec

type SecurityReference struct {
	// Reference is a reference to a security scheme defined elsewhere.
	Reference string `json:"$ref,omitempty" yaml:"$ref,omitempty"`
}
