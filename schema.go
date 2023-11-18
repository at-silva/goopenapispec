package goopenapispec

type Schema struct {
	// Type represents the type of the schema.
	Type interface{} `json:"type,omitempty" yaml:"type,omitempty"`

	// Format specifies the format of the type if applicable.
	Format string `json:"format,omitempty" yaml:"format,omitempty"`

	// Title provides a title for the schema.
	Title string `json:"title,omitempty" yaml:"title,omitempty"`

	// Description gives a description of the schema.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`

	// Default specifies the default value of the schema.
	Default interface{} `json:"default,omitempty" yaml:"default,omitempty"`

	// Enum contains a list of possible enumerated values for the schema.
	Enum []interface{} `json:"enum,omitempty" yaml:"enum,omitempty"`

	// Example holds an example value for the schema.
	Example interface{} `json:"example,omitempty" yaml:"example,omitempty"`

	// ExternalDocs points to external documentation for the schema.
	ExternalDocs *ExternalDoc `json:"externalDocs,omitempty" yaml:"externalDocs,omitempty"`

	// Items represents the schema for array types.
	Items *Schema `json:"items,omitempty" yaml:"items,omitempty"`

	// Properties contain the definitions of properties if the schema represents an object.
	Properties map[string]*Schema `json:"properties,omitempty" yaml:"properties,omitempty"`

	// AdditionalProperties represents the schema for additional properties if present.
	AdditionalProperties *Schema `json:"additionalProperties,omitempty" yaml:"additionalProperties,omitempty"`

	// Required lists the required properties for an object schema.
	Required []string `json:"required,omitempty" yaml:"required,omitempty"`

	// OneOf contains an array of subschemas, where exactly one must match.
	OneOf []Schema `json:"oneOf,omitempty" yaml:"oneOf,omitempty"`
}
