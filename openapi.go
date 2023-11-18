package goopenapispec

import (
	"encoding/json"
	"sync"
)

type (
	Paths map[string]*PathItem

	Servers map[string]*Server

	OpenAPI struct {
		// OpenAPI holds the version of the OpenAPI specification being used.
		OpenAPI string `json:"openapi" yaml:"openapi"`

		// Info contains general information about the API.
		Info Info `json:"info" yaml:"info"`

		// Paths specify the paths that can be accessed on the API.
		Paths Paths `json:"paths" yaml:"paths"`

		// Components hold reusable components that can be used in the API specification.
		Components *Components `json:"components,omitempty" yaml:"components,omitempty"`

		// Security defines security requirements that apply to the entire API.
		Security []Security `json:"security,omitempty" yaml:"security,omitempty"`

		// Servers contain network addresses for accessing the API.
		Servers Servers `json:"servers,omitempty" yaml:"servers,omitempty"`

		// ExternalDocs provide additional external documentation about the API.
		ExternalDocs *ExternalDoc `json:"externalDocs,omitempty" yaml:"externalDocs,omitempty"`
	}
)

var m sync.Mutex

// mapAddOrCreate is a utility function to add or create a value in a map based on the key.
func mapAddOrCreate[T any](mp map[string]*T, key string) (*T, bool) {
	m.Lock()
	defer m.Unlock()
	if i, ok := mp[key]; ok {
		return i, false
	}
	i := new(T)
	mp[key] = i
	return i, true
}

// AddOrCreateServer adds or creates a server based on the host URL.
func (o *OpenAPI) AddOrCreateServer(host string) *Server {
	i, c := mapAddOrCreate(o.Servers, host)
	if c {
		i.Url = host
	}

	return i
}

// AddOrCreatePathItem adds or creates a PathItem in OpenAPI Paths based on the URI.
func (o *OpenAPI) AddOrCreatePathItem(uri string) *PathItem {
	i, _ := mapAddOrCreate(o.Paths, uri)
	return i
}

// MarshalJSON converts the Servers type into JSON.
func (o Servers) MarshalJSON() ([]byte, error) {
	arr := make([]Server, 0, len(o))
	for _, v := range o {
		arr = append(arr, *v)
	}
	return json.Marshal(arr)
}
