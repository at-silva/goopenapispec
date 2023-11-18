package goopenapispec

type Contact struct {
	// Name field represents the name of the contact person or organization.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	// Url field contains the URL pointing to the contact information.
	Url string `json:"url,omitempty" yaml:"url,omitempty"`

	// Email field holds the email address of the contact person or organization.
	Email string `json:"email,omitempty" yaml:"email,omitempty"`
}
