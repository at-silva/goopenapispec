package goopenapispec

type OAuthFlows struct {
	// Implicit represents the OAuth2 implicit flow.
	Implicit OAuthFlow `json:"implicit,omitempty" yaml:"implicit,omitempty"`

	// AuthorizationCode represents the OAuth2 authorization code flow.
	AuthorizationCode OAuthFlow `json:"authorizationCode,omitempty" yaml:"authorizationCode,omitempty"`

	// Password represents the OAuth2 resource owner password flow.
	Password OAuthFlow `json:"password,omitempty" yaml:"password,omitempty"`

	// ClientCredentials represents the OAuth2 client credentials flow.
	ClientCredentials OAuthFlow `json:"clientCredentials,omitempty" yaml:"clientCredentials,omitempty"`
}
