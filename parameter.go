package goopenapispec

type Parameter struct {
	// Name is the name of the parameter.
	Name string `json:"name" yaml:"name"`

	// In is the location of the parameter. Can be "query", "header", "path", or "cookie".
	In string `json:"in" yaml:"in"`

	// Description provides a brief description of the parameter.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`

	// Required determines if the parameter is mandatory.
	Required bool `json:"required,omitempty" yaml:"required,omitempty"`

	// Deprecated specifies that the parameter is deprecated and should be avoided.
	Deprecated bool `json:"deprecated,omitempty" yaml:"deprecated,omitempty"`

	// AllowEmptyValue allows sending a parameter with an empty value.
	AllowEmptyValue bool `json:"allowEmptyValue,omitempty" yaml:"allowEmptyValue,omitempty"`

	// Style describes how the parameter value is serialized.
	Style string `json:"style,omitempty" yaml:"style,omitempty"`

	// Explode specifies behavior for parameter values of type "array" or "object".
	Explode bool `json:"explode,omitempty" yaml:"explode,omitempty"`

	// AllowReserved determines whether the parameter value should allow reserved characters.
	AllowReserved bool `json:"allowReserved,omitempty" yaml:"allowReserved,omitempty"`

	// Schema defines the type used for the parameter.
	Schema *Schema `json:"schema,omitempty" yaml:"schema,omitempty"`

	// Example provides an example value for the parameter.
	Example interface{} `json:"example,omitempty" yaml:"example,omitempty"`

	// Examples contain multiple named examples for the parameter.
	Examples map[string]*Example `json:"examples,omitempty" yaml:"examples,omitempty"`

	// Content holds representations for the parameter.
	Content map[string]*MediaType `json:"content,omitempty" yaml:"content,omitempty"`
}
