package goopenapispec

type Components struct {
	// Schemas field holds a collection of named schemas used in the OpenAPI specification.
	Schemas map[string]Schema `json:"schemas,omitempty" yaml:"schemas,omitempty"`

	// Responses field contains a collection of named responses that can be used within operations.
	Responses map[string]Response `json:"responses,omitempty" yaml:"responses,omitempty"`

	// Parameters field stores a collection of named parameters used in path, query, header, or cookie.
	Parameters map[string]Parameter `json:"parameters,omitempty" yaml:"parameters,omitempty"`

	// Examples field includes a collection of named examples that may be used in the specification.
	Examples map[string]Example `json:"examples,omitempty" yaml:"examples,omitempty"`

	// RequestBodies field contains a collection of named request bodies used in operations.
	RequestBodies map[string]RequestBody `json:"requestBodies,omitempty" yaml:"requestBodies,omitempty"`

	// Headers field holds a collection of named headers that may be used in responses.
	Headers map[string]ResponseHeader `json:"headers,omitempty" yaml:"headers,omitempty"`

	// SecuritySchemes field contains a collection of named security schemes used in the specification.
	SecuritySchemes map[string]SecurityScheme `json:"securitySchemes,omitempty" yaml:"securitySchemes,omitempty"`

	// Links field includes a collection of named links that can be used in the specification.
	Links map[string]Link `json:"links,omitempty" yaml:"links,omitempty"`

	// Callbacks field contains a collection of named callback objects used in operations.
	Callbacks map[string]Callback `json:"callbacks,omitempty" yaml:"callbacks,omitempty"`
}
