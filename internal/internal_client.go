package internal

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/valyala/fasthttp"
)

const (
	// RegisterPath 注册地址
	RegisterPath = "/api/upm-service/manage/data/register"
	// UserDetailPath 获取用户详情地址
	UserDetailPath = "/api/gateway/v1/current_user_detail"
)

type Registry struct {
	Principal string
}

type client struct {
	Options ClientOptions
}

type registryClient struct {
	Options ClientOptions
}

type Response struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data any    `json:"data"`
}

func (c *registryClient) Register(registry Registry) (*Response, error) {

	req := fasthttp.AcquireRequest()
	defer fasthttp.ReleaseRequest(req)

	req.SetRequestURI(c.Options.HostPort + RegisterPath)
	req.SetBody([]byte(registry.Principal))
	req.Header.SetContentType("application/json")
	req.Header.SetMethod("POST")
	req.Header.Add("Authorization", c.Options.AuthToken)

	resp := fasthttp.AcquireResponse()
	defer fasthttp.ReleaseResponse(resp)

	client := &fasthttp.Client{}

	if err := client.Do(req, resp); err != nil {
		return nil, err
	}

	response := &Response{}
	if err := json.Unmarshal(resp.Body(), response); err != nil {
		return nil, err
	}

	return response, nil
}

func (c *client) User(token string) (*User, error) {

	req := fasthttp.AcquireRequest()
	defer fasthttp.ReleaseRequest(req)

	req.SetRequestURI(c.Options.HostPort + UserDetailPath)
	req.Header.SetMethod("GET")
	req.Header.Add("Authorization", token)

	resp := fasthttp.AcquireResponse()
	defer fasthttp.ReleaseResponse(resp)

	client := &fasthttp.Client{}

	if err := client.Do(req, resp); err != nil {
		return nil, err
	}

	if resp.StatusCode() != fasthttp.StatusOK {
		return nil, errors.New(string(resp.Body()))
	}

	response := &Response{}

	if err := json.Unmarshal(resp.Body(), response); err != nil {
		return nil, err
	}

	if response.Code != 0 {
		return nil, errors.New(fmt.Sprintf("获取用户信息失败: %s", response.Msg))
	}

	bytes, err := json.Marshal(response.Data)

	if err != nil {
		return nil, err
	}

	user := &User{}
	err = json.Unmarshal(bytes, user)

	return user, err
}
