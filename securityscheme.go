package goopenapispec

type SecurityScheme struct {
	// Type defines the type of the security scheme, e.g., "apiKey", "http", "oauth2", or "openIdConnect".
	Type string `json:"type" yaml:"type"`

	// Description provides a short description for the security scheme.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`

	// Name specifies the name of the security scheme.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

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
}
