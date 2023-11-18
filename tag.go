package goopenapispec

type Tag struct {
	// Name represents the name of the tag and is required.
	Name string `json:"name" yaml:"name"`

	// Description provides an optional short description for the tag.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`

	// ExternalDocs contains additional external documentation for the tag.
	ExternalDocs ExternalDoc `json:"externalDocs,omitempty" yaml:"externalDocs,omitempty"`
}
