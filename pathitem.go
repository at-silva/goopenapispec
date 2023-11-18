package goopenapispec

type PathItem struct {
	// Ref is a reference to another path definition.
	Ref string `json:"$ref,omitempty" yaml:"$ref,omitempty"`

	// Summary provides a short summary of the path item.
	Summary string `json:"summary,omitempty" yaml:"summary,omitempty"`

	// Description offers a verbose explanation of the path item.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`

	// Get represents the GET operation for this path.
	Get *Operation `json:"get,omitempty" yaml:"get,omitempty"`

	// Put represents the PUT operation for this path.
	Put *Operation `json:"put,omitempty" yaml:"put,omitempty"`

	// Post represents the POST operation for this path.
	Post *Operation `json:"post,omitempty" yaml:"post,omitempty"`

	// Delete represents the DELETE operation for this path.
	Delete *Operation `json:"delete,omitempty" yaml:"delete,omitempty"`

	// Options represents the OPTIONS operation for this path.
	Options *Operation `json:"options,omitempty" yaml:"options,omitempty"`

	// Head represents the HEAD operation for this path.
	Head *Operation `json:"head,omitempty" yaml:"head,omitempty"`

	// Patch represents the PATCH operation for this path.
	Patch *Operation `json:"patch,omitempty" yaml:"patch,omitempty"`

	// Trace represents the TRACE operation for this path.
	Trace *Operation `json:"trace,omitempty" yaml:"trace,omitempty"`

	// Servers contains alternative server addresses for this path.
	Servers []Server `json:"servers,omitempty" yaml:"servers,omitempty"`

	// Parameters contain parameters applicable to all operations under this path.
	Parameters []Parameter `json:"parameters,omitempty" yaml:"parameters,omitempty"`
}
