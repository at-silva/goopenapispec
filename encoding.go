package goopenapispec

type Encoding struct {
	// ContentType field specifies the media type or content type of the encoding.
	ContentType string `json:"contentType,omitempty" yaml:"contentType,omitempty"`

	// Headers field holds a collection of headers for the encoding.
	Headers map[string]Header `json:"headers,omitempty" yaml:"headers,omitempty"`

	// Style field defines the parameter style, which affects how the parameter is serialized.
	Style ParameterStyle `json:"style,omitempty" yaml:"style,omitempty"`

	// Explode specifies whether the value should be exploded into separate values when true.
	Explode bool `json:"explode,omitempty" yaml:"explode,omitempty"`

	// AllowReserved indicates whether the parameter value can include reserved characters.
	AllowReserved bool `json:"allowReserved,omitempty" yaml:"allowReserved,omitempty"`
}
