package goopenapispec

type OAuthFlow struct {
	// AuthorizationURL represents the URL where the user can be redirected to authorize the application.
	AuthorizationURL string `json:"authorizationUrl,omitempty" yaml:"authorizationUrl,omitempty"`

	// TokenURL is the URL for the server's token endpoint in the OAuth flow.
	TokenURL string `json:"tokenUrl,omitempty" yaml:"tokenUrl,omitempty"`

	// RefreshURL specifies the URL where the token can be refreshed in the OAuth flow.
	RefreshURL string `json:"refreshUrl,omitempty" yaml:"refreshUrl,omitempty"`

	// Scopes contain a map of available scopes for the OAuth flow.
	Scopes map[string]string `json:"scopes,omitempty" yaml:"scopes,omitempty"`
}
