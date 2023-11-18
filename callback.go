package goopenapispec

type Callback struct {
	// A map of callback URLs to their respective RequestBodies.
	// The key is the name of the callback URL, and the value is a `RequestBody` object.
	CallbackItems map[string]RequestBody `json:"-" yaml:"-"`

	// This field is used to document the callbacks.
	// It should be a map where the keys are the names of the callback events,
	// and the values are `PathItem` objects that describe the requests that are sent to the callback URLs.
	// Note that the `Operation` object inside the `PathItem` should have an empty `Responses` object,
	// since the response schema is defined in the `RequestBody` object in the `CallbackItems` map.
	CallbackPaths Paths `json:"-" yaml:"-"`
}
		