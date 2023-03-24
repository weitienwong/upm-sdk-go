package auth

import (
	"github.com/weitienwong/upm-sdk-go/pkg/client"
	"github.com/weitienwong/upm-sdk-go/pkg/match"
	"log"
)

func Authorize(path, method, token string, client client.Client) bool {
	user, err := client.User(token)
	if err != nil {
		log.Println("获取UPM用户信息失败\n", err)
		return false
	}
	// 遍历应用
	for _, app := range user.Apps {
		var patterns []string
		// 遍历授权API
		for _, api := range app.Apis {
			// 过滤请求方法与参数method相同的授权API
			if method == api.RequestMethod {
				patterns = append(patterns, api.Url)
			}
		}
		// 是否命中
		if match.HitMatch(path, patterns) {
			return true
		}
	}
	return false
}
