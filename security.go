package goopenapispec

type Security struct {
	// Name specifies the name of the security scheme.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	// Type defines the type of the security scheme, such as "apiKey", "http", "oauth2", or "openIdConnect".
	Type string `json:"type,omitempty" yaml:"type,omitempty"`

	// Description provides a short description for the security scheme.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`

	// In determines the location of the API key or bearer token, e.g., "header", "query", or "cookie".
	In string `json:"in,omitempty" yaml:"in,omitempty"`

	// Scheme specifies the name of the HTTP authentication scheme to be used, e.g., "basic", "bearer", or "digest".
	Scheme string `json:"scheme,omitempty" yaml:"scheme,omitempty"`

	// BearerFormat defines the format of the bearer token, e.g., "JWT".
	BearerFormat string `json:"bearerFormat,omitempty" yaml:"bearerFormat,omitempty"`

	// Flows contain available OAuth flows for the security scheme.
	Flows OAuthFlows `json:"flows,omitempty" yaml:"flows,omitempty"`

	// OpenIDConnectURL holds the URL to the OpenID Connect discovery endpoint.
	OpenIDConnectURL string `json:"openIdConnectUrl,omitempty" yaml:"openIdConnectUrl,omitempty"`

	// Reference is a reference to a security scheme defined elsewhere.
	Reference SecurityReference `json:"$ref,omitempty" yaml:"$ref,omitempty"`
}
