package client

import "github.com/weitienwong/upm-sdk-go/internal"

type (
	// Options are optional parameters for Client creation
	Options = internal.ClientOptions
	// Registry is principal of authority resource
	Registry = internal.Registry
	// Response is the response from UPM
	Response = internal.Response
	// User is the structure of user from UPM
	User = internal.User

	// Client is the client for getting information about user
	Client interface {
		// User request user information from UPM
		User(token string) (*User, error)
	}
	// RegistryClient is the client for managing the resource of authority
	RegistryClient interface {
		// Register resources with UPM server
		Register(registry Registry) (*Response, error)
	}
)

// NewClient create a instance of Client
func NewClient(options Options) (Client, error) {
	return internal.NewClient(options)
}

// NewRegistryClient create a instance of RegistryClient
func NewRegistryClient(options Options) (RegistryClient, error) {
	return internal.NewRegistryClient(options)
}
