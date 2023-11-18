package goopenapispec

type Response struct {
	// Description provides a description of the response.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`

	// Headers represent the headers the response can have.
	Headers map[string]ResponseHeader `json:"headers,omitempty" yaml:"headers,omitempty"`

	// Content contains media types the response can have.
	Content map[string]MediaType `json:"content,omitempty" yaml:"content,omitempty"`

	// Links contain links the response can have.
	Links map[string]Link `json:"links,omitempty" yaml:"links,omitempty"`
}
