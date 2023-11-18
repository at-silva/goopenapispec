package goopenapispec

type Link struct {
	// Ref is a reference to another link definition.
	Ref string `json:"$ref,omitempty" yaml:"$ref,omitempty"`

	// OperationID specifies the ID of the linked operation in the OpenAPI document.
	OperationID string `json:"operationId,omitempty" yaml:"operationId,omitempty"`

	// Parameters contains information about parameters linked to the linked operation.
	// It can be a map[string]interface{} or an array, depending on the structure.
	Parameters interface{} `json:"parameters,omitempty" yaml:"parameters,omitempty"`

	// RequestBody contains information about the request body linked to the linked operation.
	// It can be a map[string]interface{} or an array, depending on the structure.
	RequestBody interface{} `json:"requestBody,omitempty" yaml:"requestBody,omitempty"`

	// Description provides additional details about the purpose or functionality of the link.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`

	// Server specifies the server URL or details for the linked operation.
	Server Server `json:"server,omitempty" yaml:"server,omitempty"`
}
