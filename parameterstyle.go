package goopenapispec

type ParameterStyle struct {
	// Matrix specifies the "matrix" style for parameter serialization.
	Matrix *struct{} `json:"matrix,omitempty" yaml:"matrix,omitempty"`

	// Label specifies the "label" style for parameter serialization.
	Label *struct{} `json:"label,omitempty" yaml:"label,omitempty"`

	// Form specifies the "form" style for parameter serialization.
	Form *struct{} `json:"form,omitempty" yaml:"form,omitempty"`

	// Simple specifies the "simple" style for parameter serialization.
	Simple *struct{} `json:"simple,omitempty" yaml:"simple,omitempty"`

	// SpaceDelimited specifies the "spaceDelimited" style for parameter serialization.
	SpaceDelimited *struct{} `json:"spaceDelimited,omitempty" yaml:"spaceDelimited,omitempty"`

	// PipeDelimited specifies the "pipeDelimited" style for parameter serialization.
	PipeDelimited *struct{} `json:"pipeDelimited,omitempty" yaml:"pipeDelimited,omitempty"`

	// DeepObject specifies the "deepObject" style for parameter serialization.
	DeepObject *struct{} `json:"deepObject,omitempty" yaml:"deepObject,omitempty"`
}
