package test

import (
	"encoding/json"
	"fmt"
	"github.com/weitienwong/upm-sdk-go/pkg/upm"
	"testing"
)

const principal = `{
	"serviceKey": "dataplat-xxx-001",
	"serviceAccess": null,
	"apps": [
		{
			"appCode": "chaos-engineering-app-01",
			"appTitle": "混沌工程-01",
			"appUrl": "http://cloudwrecker.test.bizseer.com/",
			"appIcon": null,
			"groupName": "平台管理",
			"groupIcon": "platform-management",
			"groupPanelName": "数据中台",
			"reqApis": [
				{
					"httpMethod": "GET",
					"path":"/cloudbed/static/{filepath:*}",
					"name":"静态资源"
				}
			],
			"menus": [
				{
					"code": "CW-001",
					"title": "稳定性平台",
					"path": "",
					"icon": "storage",
					"subMenus": [
						{
							"code": "CW-001-001",
							"title": "概览",
							"path": "/fault-injection/overview",
							"icon": "storage",
							"subMenus": []
						}, {
							"code": "CW-001-002",
							"title": "场景管理",
							"path": "/fault-injection/experiment",
							"icon": "storage",
							"subMenus": []
						}, {
							"code": "CW-001-003",
							"title": "专家经验库",
							"path": "/fault-injection/expert-library/",
							"icon": "storage",
							"subMenus": []
						}, {
							"code": "CW-001-004",
							"title": "运维对抗",
							"path": "/fault-injection/batch",
							"icon": "storage",
							"subMenus": []
						}, {
							"code": "CW-001-005",
							"title": "定时计划",
							"path": "/fault-injection/schedule",
							"icon": "storage",
							"subMenus": []
						}, {
							"code": "CW-001-006",
							"title": "历史实验",
							"path": "/fault-injection/experiment-history",
							"icon": "storage",
							"subMenus": []
						}, {
							"code": "CW-001-007",
							"title": "爆破点管理",
							"path": "/fault-injection/shot-point",
							"icon": "storage",
							"subMenus": []
						}
					]
				}
			],
			"menuRegisterStrategy": 0,
			"buttons": [
				{
					"key": "dm_source_delete",
					"title": "删除按钮"
				},
				{
					"key": "dm_source_update",
					"title": "更新按钮"
				}
			],
			"buttonRegisterStrategy": null,
			"resourceBuckets": [
				{
					"name": "权限场景003",
					"apiMatches": [
						{
							"reqMethods": ["GET"],
							"pattern": "/cloudbed/static/{filepath:*}"
						}, {
							"reqMethods": ["GET","POST"],
							"pattern": "/cloudbed/schedule"
						}, {
							"reqMethods": ["GET"],
							"pattern": "/cloudbed/schedule/unexpired"
						}, {
							"reqMethods": ["GET","POST","DELETE"],
							"pattern": "/cloudbed/schedule/{schedule_id:*}"
						}
					],
					"menuCodes": [
						"data-source-001"
					],
					"menuPaths": null,
					"buttonCodes": [
						"dm_source_delete"
					]
				}
			],
			"resourceBucketRegisterStrategy": null,
			"dataRestrictions": [],
			"dataRestrictionRegisterStrategy": null
		}
	]
}`

func TestRegistry(t *testing.T) {
	c, err := upm.NewRegistryClient(upm.ClientOptions{
		HostPort:  "http://gateway.dev.bizseer.com",
		AuthToken: "Bearer eyJhbGciOiJIUzUxMiJ9.eyJleHAiOjE2ODAyNTE2MDIsInN1YiI6ImxpeWFuZ3lhbmciLCJjcmVhdGVkIjoxNjc5NjQ2ODAyMDgzfQ.OVCSeMreKDPaCSdiRl7z9gqpRSupzxK9JPcg8UL04J1oEGHvc_aAQJdcLHFcjsEWtmJDOI1B6mNsSlwxU0ZU7A",
	})

	if err != nil {
		t.Error("无法建立连接")
	}

	resp, err := c.Register(upm.Registry{Principal: principal})
	if err != nil {
		t.Error(err)
	}
	fmt.Println(resp)
}

func TestUserInfo(t *testing.T) {
	c, err := upm.NewClient(upm.ClientOptions{
		HostPort: "http://gateway.dev.bizseer.com",
	})

	if err != nil {
		t.Error("无法建立连接")
	}

	token := "Bearer eyJhbGciOiJIUzUxMiJ9.eyJleHAiOjE2ODAyNTE2MDIsInN1YiI6ImxpeWFuZ3lhbmciLCJjcmVhdGVkIjoxNjc5NjQ2ODAyMDgzfQ.OVCSeMreKDPaCSdiRl7z9gqpRSupzxK9JPcg8UL04J1oEGHvc_aAQJdcLHFcjsEWtmJDOI1B6mNsSlwxU0ZU7A"

	u, err := c.User(token)

	if err != nil {
		t.Error("获取用户失败")
	}

	b, _ := json.Marshal(u)
	fmt.Println(string(b))
}
