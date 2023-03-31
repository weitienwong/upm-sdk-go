package test

import (
	"encoding/json"
	"fmt"
	"github.com/weitienwong/upm-sdk-go/pkg/upm"
	"io/ioutil"
	"testing"
)

const token = "Bearer eyJhbGciOiJIUzUxMiJ9.eyJleHAiOjE2ODA3NjYxNzgsInN1YiI6ImFkbWluIiwiY3JlYXRlZCI6MTY4MDE2MTM3ODc1OH0.hOgZm9S4ejERER1U4eb1h5tdlgXhVw5dCylUd6CadUDdrZxwu3u7r2zN1jnLiXWsvLpiINMDpm0icHSIa2TthQ"

// 注册
func TestRegistry(t *testing.T) {
	bytes, err := ioutil.ReadFile("./upm.json")
	if err != nil {
		fmt.Println(err)
		return
	}
	c, err := upm.NewRegistryClient(upm.ClientOptions{
		HostPort:  "http://gateway.dev.bizseer.com",
		AuthToken: token,
	})

	if err != nil {
		t.Error("无法建立连接")
	}

	resp, err := c.Register(upm.Registry{Principal: string(bytes)})
	if err != nil {
		t.Error(err)
	}
	fmt.Println(resp)
}

// 获取用户信息
func TestUserInfo(t *testing.T) {
	c, err := upm.NewClient(upm.ClientOptions{
		HostPort: "http://gateway.dev.bizseer.com",
	})

	if err != nil {
		t.Error("无法建立连接")
	}

	//token := "Bearer eyJhbGciOiJIUzUxMiJ9.eyJleHAiOjE2ODAyNTE2MDIsInN1YiI6ImxpeWFuZ3lhbmciLCJjcmVhdGVkIjoxNjc5NjQ2ODAyMDgzfQ.OVCSeMreKDPaCSdiRl7z9gqpRSupzxK9JPcg8UL04J1oEGHvc_aAQJdcLHFcjsEWtmJDOI1B6mNsSlwxU0ZU7A"

	u, err := c.User(token)

	if err != nil {
		t.Error("获取用户失败")
	}

	b, _ := json.Marshal(u)
	fmt.Println(string(b))
}
