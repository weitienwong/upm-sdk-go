package auth

import (
	"github.com/weitienwong/upm-sdk-go/pkg/client"
	"github.com/weitienwong/upm-sdk-go/pkg/match"
	"log"
)

func Authorize(path, token string, client client.Client) bool {
	_, err := client.User(token)
	if err != nil {
		log.Println("获取UPM用户信息失败\n", err)
		return false
	}
	// 遍历用户权限是否匹配path
	patterns := []string{}
	return match.HitMatch(path, patterns)
}
