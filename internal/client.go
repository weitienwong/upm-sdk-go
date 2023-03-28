package internal

import (
	"errors"
	"fmt"
	"net"
	"strings"
	"time"
)

// Client UPM客户端接口
type Client interface {
	User(token string) (*User, error)
}

// RegistryClient 注册资源客户端
type RegistryClient interface {
	Register(registry Registry) (*Response, error)
}

// ClientOptions configuration parameters for creating a Client or RegistryClient.
type ClientOptions struct {
	HostPort  string
	AuthToken string
}

// NewClient creates an instance of an upm Client
func NewClient(options ClientOptions) (Client, error) {
	err := dial(options.HostPort)

	if err != nil {
		return nil, err
	}
	return &client{
		Options: options,
	}, nil
}

// NewRegistryClient creates an instance of an upm registry client
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

	var protocal string

	if strings.HasPrefix(hostPort, "http") {
		protocal = "http"
	}
	if strings.HasPrefix(hostPort, "https") {
		protocal = "https"
	}

	if protocal != "" {
		hostPort = strings.TrimPrefix(hostPort, protocal+"://")
	}

	arr := strings.Split(hostPort, ":")

	if len(arr) == 2 {
		hostPort = arr[0]
		protocal = arr[1]
	}
	hostPort = fmt.Sprintf("%s:%s", hostPort, protocal)
	return dialTimeout(hostPort, time.Second)
}

// connects to the address on the named network taking a timeout.
func dialTimeout(hostport string, timeout time.Duration) error {
	if timeout == 0 {
		timeout = time.Second
	}
	_, err := net.DialTimeout("tcp", hostport, timeout)

	return err
}
