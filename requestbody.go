package goopenapispec

type RequestBody struct {
	// Description provides a brief description of the request body.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`

	// Content contains expected media types and their corresponding schema for the request body.
	Content map[string]MediaType `json:"content" yaml:"content"`

	// Required determines whether the request body is required. Default value is false.
	Required bool `json:"required,omitempty" yaml:"required,omitempty"`
}
