package goopenapispec

type Example struct {
	// Summary field represents a brief summary or title for the example.
	Summary string `json:"summary,omitempty" yaml:"summary,omitempty"`

	// Description field contains a longer description providing more details about the example.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`

	// Value field holds the actual value of the example, which can be of any type.
	Value interface{} `json:"value,omitempty" yaml:"value,omitempty"`

	// ExternalValue field references an external resource for the example's value.
	ExternalValue string `json:"externalValue,omitempty" yaml:"externalValue,omitempty"`
}
