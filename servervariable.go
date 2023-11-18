package goopenapispec

type ServerVariable struct {
	// Enum is an array of possible values for the variable.
	Enum []string `json:"enum,omitempty" yaml:"enum,omitempty"`

	// Default represents the default value of the variable and is required.
	Default string `json:"default" yaml:"default"`

	// Description provides an optional description of the variable.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`
}
