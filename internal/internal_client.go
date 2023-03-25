package internal

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/valyala/fasthttp"
)

const RegisterPath = "/api/upm-service/manage/data/register"
const UserInfoPath = "/api/gateway/v1/current_user_detail"

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

func (r *registryClient) Register(registry Registry) (*Response, error) {

	req := fasthttp.AcquireRequest()
	defer fasthttp.ReleaseRequest(req)

	uri := r.Options.HostPort + RegisterPath
	req.SetRequestURI(uri)

	body := []byte(registry.Principal)
	req.SetBody(body)

	req.Header.SetContentType("application/json")
	req.Header.SetMethod("POST")
	req.Header.Add("Authorization", r.Options.AuthToken)

	resp := fasthttp.AcquireResponse()

	defer fasthttp.ReleaseResponse(resp)
	client := &fasthttp.Client{}

	if err := client.Do(req, resp); err != nil {
		return nil, err
	}

	response := &Response{}
	err := json.Unmarshal(resp.Body(), response)

	if err != nil {
		return nil, err
	}
	return response, nil
}

func (c *client) User(token string) (*User, error) {

	req := fasthttp.AcquireRequest()
	defer fasthttp.ReleaseRequest(req)

	uri := c.Options.HostPort + UserInfoPath
	req.SetRequestURI(uri)
	req.Header.SetMethod("GET")
	req.Header.Add("Authorization", token)

	resp := fasthttp.AcquireResponse()

	defer fasthttp.ReleaseResponse(resp)
	client := &fasthttp.Client{}

	if err := client.Do(req, resp); err != nil {
		return nil, err
	}

	response := &Response{}
	err := json.Unmarshal(resp.Body(), response)
	if err != nil {
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
