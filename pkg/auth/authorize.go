package auth

import (
	"github.com/weitienwong/upm-sdk-go/pkg/match"
	"github.com/weitienwong/upm-sdk-go/pkg/upm"
	"strings"
)

// Authorize 根据path和method匹配用户授权的API资源，命中则返回true，否则返回false
func Authorize(path, method string, apps []*upm.App) bool {
	if len(apps) == 0 {
		return false
	}
	// 遍历应用
	for _, app := range apps {
		var patterns []string
		// 遍历授权API
		for _, api := range app.Apis {
			// 过滤请求方法与参数method相同的授权API
			if strings.EqualFold(method, api.RequestMethod) {
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
