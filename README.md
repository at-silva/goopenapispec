# GoOpenAPISpec Package

The `goopenapispec` package provides a limited set of structures and methods to generate OpenAPI specifications in Go. It allows you to create, modify, and manage OpenAPI specs files programmatically.

## Overview

The OpenAPI structure forms the core of the package, encompassing various elements of an OpenAPI specification:

- **Info:** General information about the API.
- **Paths**: Defines the available paths and operations on the API.
- **Components**: Reusable components used in the API specification.
- **Security**: Security requirements for the API.
- **Servers**: Network addresses for accessing the API.
- **ExternalDocs**: Additional external documentation about the API.

## Features

- **PathItem Operations**: Create or modify operations (GET, POST, PUT, DELETE, etc.) for specific paths.
- **Responses**: Manage and define possible responses for operations.
- **Security** Schemes: Define security schemes and requirements for API endpoints.
    Components: Reusable components such as schemas, parameters, request bodies, responses, etc.
- **Servers**: Define and manage network addresses for API access.

## Getting Started

To begin using the openapispec package:

**Installation**: Install the package by importing it into your Go project.

```shell
go get github.com/at-silva/goopenapispec
```
*Usage*: Import the package into your Go code and start interacting with OpenAPI specifications programmatically.

Here's a simple example:

```go
package main

import (
	"github.com/at-silva/goopenapispec"
)

func main() {
	// Create a new OpenAPI specification
	api := openapispec.OpenAPI{
		OpenAPI: "3.0.0",
		Info: openapispec.Info{
			Title:   "My API",
			Version: "1.0.0",
			// ... other fields
		},
		// ... other components
	}

	// Modify or use various elements of the OpenAPI specification using package methods.
}
```

## Contribution Guidelines

We welcome contributions! To contribute to the openapispec package, follow these steps:

1. Fork the repository.
2. Make changes in your fork.
3. Submit a pull request with a clear description of the changes made.

## License

This package is licensed under the MIT License - see the LICENSE file for details.

Feel free to add more sections such as "Examples", "Advanced Usage", or "Contributing Guidelines" based on the specifics of your package and its usage. Adjust the sections and details as needed to best represent your package on GitHub!