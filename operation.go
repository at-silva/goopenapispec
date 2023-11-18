package goopenapispec

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

var caser = cases.Title(language.AmericanEnglish)

type Operation struct {
	// Tags are used for API documentation control, providing logical grouping of operations.
	Tags []string `json:"tags,omitempty" yaml:"tags,omitempty"`

	// Summary provides a short summary of the operation's purpose.
	Summary string `json:"summary,omitempty" yaml:"summary,omitempty"`

	// Description offers a verbose explanation of the operation's behavior.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`

	// ExternalDocs holds additional external documentation for this operation.
	ExternalDocs ExternalDoc `json:"externalDocs,omitempty" yaml:"externalDocs,omitempty"`

	// OperationID is a unique string used to identify the operation.
	OperationID string `json:"operationId,omitempty" yaml:"operationId,omitempty"`

	// Parameters specify the parameters needed to send to the operation.
	Parameters []Parameter `json:"parameters,omitempty" yaml:"parameters,omitempty"`

	// RequestBody represents the request body needed to send to the operation.
	RequestBody *RequestBody `json:"requestBody,omitempty" yaml:"requestBody,omitempty"`

	// Responses contain possible responses returned by the operation.
	Responses map[string]*ResponseRef `json:"responses" yaml:"responses"`

	// Deprecated declares if this operation is deprecated.
	Deprecated bool `json:"deprecated,omitempty" yaml:"deprecated,omitempty"`

	// Security declares which security mechanisms can be used for this operation.
	Security []Security `json:"security,omitempty" yaml:"security,omitempty"`

	// Servers offer an alternative server array to service this operation.
	Servers []Server `json:"servers,omitempty" yaml:"servers,omitempty"`
}

// newOperation creates a new Operation based on URL and method.
func newOperation(url, method string) *Operation {
	op := new(Operation)
	op.Responses = map[string]*ResponseRef{}

	s := strings.Split(url, "/")
	for i := range s {
		s[i] = caser.String(s[i])
	}

	op.OperationID = strings.ToLower(strings.TrimSpace(method)) + strings.Join(s, "")

	return op
}

// AddOrCreateOperation adds or creates an Operation in a PathItem based on URL and method.
func (p *PathItem) AddOrCreateOperation(url, method string) (*Operation, error) {
	m.Lock()
	defer m.Unlock()
	switch strings.TrimSpace(strings.ToUpper(method)) {
	case http.MethodPost:
		if p.Post == nil {
			p.Post = newOperation(url, method)
		}
		return p.Post, nil
	case http.MethodGet:
		if p.Get == nil {
			p.Get = newOperation(url, method)
		}
		return p.Get, nil
	case http.MethodPatch:
		if p.Patch == nil {
			p.Patch = newOperation(url, method)
		}
		return p.Patch, nil
	case http.MethodPut:
		if p.Put == nil {
			p.Put = newOperation(url, method)
		}
		return p.Put, nil
	case http.MethodDelete:
		if p.Delete == nil {
			p.Delete = newOperation(url, method)
		}
		return p.Delete, nil
	case http.MethodOptions:
		if p.Options == nil {
			p.Options = newOperation(url, method)
		}
		return p.Options, nil
	case http.MethodHead:
		if p.Head == nil {
			p.Head = newOperation(url, method)
		}
		return p.Head, nil
	case http.MethodTrace:
		if p.Trace == nil {
			p.Trace = newOperation(url, method)
		}
		return p.Trace, nil
	}

	return nil, fmt.Errorf("invalid method: %s", method)
}

// AddOrCreateResponseRef adds or creates a ResponseRef in an Operation based on status code.
func (o *Operation) AddOrCreateResponseRef(status int) *ResponseRef {
	r, c := mapAddOrCreate(o.Responses, strconv.Itoa(status))
	if c {
		r.Description = http.StatusText(status)
	}

	return r
}
