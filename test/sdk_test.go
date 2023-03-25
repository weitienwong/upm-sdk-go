package test

import (
	"encoding/json"
	"fmt"
	"github.com/weitienwong/upm-sdk-go/pkg/upm"
	"testing"
)

const principal = `{
	"serviceKey": "chaos-engineering-app-02",
	"serviceAccess": null,
	"apps": [
		{
			"appCode": "chaos-engineering-app-02",
			"appTitle": "混沌工程-02",
			"appUrl": "http://cloudwrecker.test.bizseer.com/",
			"appIcon": null,
			"groupName": "混沌工程",
			"groupIcon": "platform-management",
			"groupPanelName": "混沌工程",
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
					"title": "概览",
					"path": "/overview",
					"icon": "storage",
					"subMenus": []
				}, {
					"code": "CW-002",
					"title": "场景管理",
					"path": "/experiment",
					"icon": "storage",
					"subMenus": []
				}, {
					"code": "CW-003",
					"title": "专家经验库",
					"path": "/expert-library",
					"icon": "storage",
					"subMenus": []
				}, {
					"code": "CW-004",
					"title": "运维对抗",
					"path": "/batch",
					"icon": "storage",
					"subMenus": []
				}, {
					"code": "CW-005",
					"title": "定时计划",
					"path": "/schedule",
					"icon": "storage",
					"subMenus": []
				}, {
					"code": "CW-006",
					"title": "历史实验",
					"path": "/experiment-history",
					"icon": "storage",
					"subMenus": []
				}, {
					"code": "CW-007",
					"title": "爆破点管理",
					"path": "/shot-point",
					"icon": "storage",
					"subMenus": []
				}, {
					"code": "CW-008",
					"title": "能力模板管理",
					"path": "/attack-template",
					"icon": "storage",
					"subMenus": []
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
const token = "Bearer eyJhbGciOiJIUzUxMiJ9.eyJleHAiOjE2ODAwNzQ2NDcsInN1YiI6ImFkbWluIiwiY3JlYXRlZCI6MTY3OTQ2OTg0Nzc2MH0.nvwjw60C6e-e7xm09uPnEXZLzODtbH0MzJR_eDCmPzjMbHoz3O2gnIJO4rVbVpIn_SDzb7VBYmVQdpJ-FTP0_w"

func TestRegistry(t *testing.T) {
	c, err := upm.NewRegistryClient(upm.ClientOptions{
		HostPort:  "http://10.0.100.244:38086",
		AuthToken: token,
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
