package goopenapispec

type Header struct {
	// Description provides additional information about the header.
	Description string `json:"description,omitempty"`

	// Required indicates whether the header is mandatory.
	Required bool `json:"required,omitempty"`

	// Deprecated specifies if the header is deprecated or no longer in use.
	Deprecated bool `json:"deprecated,omitempty"`

	// Explode specifies whether the header value should be exploded into multiple values.
	Explode bool `json:"explode,omitempty"`

	// AllowEmptyValue indicates if the header allows an empty value.
	AllowEmptyValue bool `json:"allowEmptyValue,omitempty"`

	// AllowReserved specifies whether reserved characters are allowed in the header value.
	AllowReserved bool `json:"allowReserved,omitempty"`

	// Style describes how the header value should be serialized.
	Style string `json:"style,omitempty"`

	// Schema contains the schema defining the structure of the header.
	Schema *Schema `json:"schema,omitempty"`

	// Example holds an example value for the header.
	Example interface{} `json:"example,omitempty"`

	// Examples holds multiple named examples for the header.
	Examples map[string]*Example `json:"examples,omitempty"`

	// Content contains media types for the header's content.
	Content map[string]*MediaType `json:"content,omitempty"`
}
