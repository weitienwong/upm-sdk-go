package upm

import "github.com/weitienwong/upm-sdk-go/internal"

type (
	// ClientOptions are optional parameters for Client creation
	ClientOptions = internal.ClientOptions
	// Registry is principal of authority resource
	Registry = internal.Registry
	// Response is the response from UPM
	Response = internal.Response
	// UserDetail is the structure of user from UPM
	UserDetail = internal.User
	// App a set of resources
	App = internal.App

	// Client is the client for getting information about user
	Client interface {
		// GetUserDetail get user information from UPM using the token
		GetUserDetail(token string) (*UserDetail, error)
		// GetUserDetailBySecret get user information from UPM using the key and secret
		GetUserDetailBySecret(key, secret string) (*UserDetail, error)
	}
	// RegistryClient is the client for managing the resource of authority
	RegistryClient interface {
		// Register resources with UPM server
		Register(registry Registry) (*Response, error)
	}
)

// NewClient create an instance of Client
func NewClient(options ClientOptions) (Client, error) {
	return internal.NewClient(options)
}

// NewRegistryClient create an instance of RegistryClient
func NewRegistryClient(options ClientOptions) (RegistryClient, error) {
	return internal.NewRegistryClient(options)
}
