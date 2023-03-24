package internal

import (
	"errors"
	"net"
	"strings"
	"time"
)

type Client interface {
	User(token string) (*User, error)
}

type RegistryClient interface {
	Register(registry Registry) (*Response, error)
}

type ClientOptions struct {
	HostPort  string
	AuthToken string
}

func NewClient(options ClientOptions) (Client, error) {
	err := dial(options.HostPort)

	if err != nil {
		return nil, err
	}
	return &client{
		Options: options,
	}, nil
}

func NewRegistryClient(options ClientOptions) (RegistryClient, error) {

	if options.AuthToken == "" {
		return nil, errors.New("authToken not provide")
	}

	err := dial(options.HostPort)

	if err != nil {
		return nil, err
	}

	return &registryClient{
		options,
	}, nil
}

// connects to the address on the named network
func dial(hostPort string) error {
	if hostPort == "" {
		return errors.New("hostPort not provide")
	}
	return dialTimeout(strings.TrimPrefix(hostPort, "http://"), time.Second)
}

// connects to the address on the named network taking a timeout.
func dialTimeout(hostport string, timeout time.Duration) error {
	if timeout == 0 {
		timeout = time.Second
	}
	_, err := net.DialTimeout("tcp", hostport, timeout)
	return err
}
