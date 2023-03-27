package test

import (
	"encoding/json"
	"fmt"
	"github.com/weitienwong/upm-sdk-go/pkg/upm"
	"io/ioutil"
	"testing"
)

const principal = `{
    "serviceKey": "cloudwrecker",
    "serviceAccess": null,
    "apps": [
        {
            "appCode": "cloudwrecker",
            "appTitle": "稳定性平台",
            "appUrl": "http://cloudwrecker.test.bizseer.com/",
            "appIcon": null,
            "groupName": "",
            "groupIcon": "platform-management",
            "groupPanelName": "",
            "reqApis": [
                {
                    "httpMethod": "GET",
                    "path": "/cloudbed/static/*",
                    "name": "静态资源"
                },
                {
                    "httpMethod": "GET",
                    "path": "/cloudbed/schedule",
                    "name": "定时任务分页"
                },
                {
                    "httpMethod": "POST",
                    "path": "/cloudbed/schedule",
                    "name": "创建定时任务"
                },
                {
                    "httpMethod": "GET",
                    "path": "/cloudbed/schedule/unexpired",
                    "name": "获取未过期定时任务"
                },
                {
                    "httpMethod": "GET",
                    "path": "/cloudbed/schedule/*",
                    "name": "获取定时任务"
                },
                {
                    "httpMethod": "POST",
                    "path": "/cloudbed/schedule/*",
                    "name": "修改定时任务"
                },
                {
                    "httpMethod": "DELETE",
                    "path": "/cloudbed/schedule/*",
                    "name": "删除定时任务"
                },
                {
                    "httpMethod": "GET",
                    "path": "/cloudbed/statistics/aggr",
                    "name": "概览"
                },
                {
                    "httpMethod": "GET",
                    "path": "/cloudbed/batch",
                    "name": "获取运维对抗"
                },
                {
                    "httpMethod": "POST",
                    "path": "/cloudbed/batch",
                    "name": "创建运维对抗"
                },
                {
                    "httpMethod": "PUT",
                    "path": "/cloudbed/batch",
                    "name": "修改运维对抗"
                },
                {
                    "httpMethod": "POST",
                    "path": "/cloudbed/batch/verify",
                    "name": "运维对抗校验"
                },
                {
                    "httpMethod": "GET",
                    "path": "/cloudbed/setting/*",
                    "name": "根据类型获取全局配置"
                },
                {
                    "httpMethod": "PUT",
                    "path": "/cloudbed/setting/*",
                    "name": "修改全局配置"
                },
                {
                    "httpMethod": "GET",
                    "path": "/cloudbed/asset/host",
                    "name": "爆破点(主机)"
                },
                {
                    "httpMethod": "POST",
                    "path": "/cloudbed/asset/host",
                    "name": "创建爆破点(主机)"
                },
                {
                    "httpMethod": "POST",
                    "path": "/cloudbed/agent/registry/host",
                    "name": "注册爆破点"
                },
                {
                    "httpMethod": "POST",
                    "path": "/cloudbed/agent/refresh/host",
                    "name": "刷新爆破点"
                },
                {
                    "httpMethod": "GET",
                    "path": "/cloudbed/asset/host/tag",
                    "name": "爆破点标签列表"
                },
                {
                    "httpMethod": "GET",
                    "path": "/cloudbed/asset/host/*",
                    "name": "根据ID获取爆破点"
                },
                {
                    "httpMethod": "PUT",
                    "path": "/cloudbed/asset/host/*",
                    "name": "修改爆破点"
                },
                {
                    "httpMethod": "DELETE",
                    "path": "/cloudbed/asset/host/*",
                    "name": "删除爆破点"
                },
                {
                    "httpMethod": "GET",
                    "path": "/cloudbed/asset/kubernetes",
                    "name": "爆破点列表(k8s)"
                },
                {
                    "httpMethod": "POST",
                    "path": "/cloudbed/asset/kubernetes",
                    "name": "创建爆破点(k8s)"
                },
                {
                    "httpMethod": "GET",
                    "path": "/cloudbed/asset/kubernetes/*",
                    "name": "根据ID获取爆破点(k8s)"
                },
                {
                    "httpMethod": "PUT",
                    "path": "/cloudbed/asset/kubernetes/*",
                    "name": "修改爆破点(k8s)"
                },
                {
                    "httpMethod": "DELETE",
                    "path": "/cloudbed/asset/kubernetes/*",
                    "name": "删除爆破点(k8s)"
                },
                {
                    "httpMethod": "GET",
                    "path": "/cloudbed/asset/kubernetes/*/ns",
                    "name": "获取命名空间"
                },
                {
                    "httpMethod": "GET",
                    "path": "/cloudbed/asset/kubernetes/*/*/pod",
                    "name": "获取POD"
                },
                {
                    "httpMethod": "GET",
                    "path": "/cloudbed/asset/kubernetes/*/*/tag",
                    "name": "获取命名空间POD"
                },
                {
                    "httpMethod": "POST",
                    "path": "/cloudbed/inspection/test",
                    "name": "状态检查(测试)"
                },
                {
                    "httpMethod": "POST",
                    "path": "/cloudbed/tag",
                    "name": "添加标签"
                },
                {
                    "httpMethod": "GET",
                    "path": "/cloudbed/tag/all",
                    "name": "获取所有标签"
                },
                {
                    "httpMethod": "GET",
                    "path": "/cloudbed/prometheus/label/*",
                    "name": "过滤Prometheus标签"
                },
                {
                    "httpMethod": "GET",
                    "path": "/cloudbed/indication/data",
                    "name": "获取稳态监控数据"
                },
                {
                    "httpMethod": "POST",
                    "path": "/cloudbed/temporal/namespace/*",
                    "name": "创建Temporal命名空间"
                },
                {
                    "httpMethod": "POST",
                    "path": "/cloudbed/scene/category",
                    "name": "添加原子能力类目"
                },
                {
                    "httpMethod": "PUT",
                    "path": "/cloudbed/scene/category",
                    "name": "修改原子能力类目"
                },
                {
                    "httpMethod": "DELETE",
                    "path": "/cloudbed/scene/category/*",
                    "name": "删除原子能力类目"
                },
                {
                    "httpMethod": "GET",
                    "path": "/cloudbed/scene/category/*",
                    "name": "获取原子能力类目"
                },
                {
                    "httpMethod": "GET",
                    "path": "/cloudbed/scene/category/all",
                    "name": "获取原子能力所有类目"
                },
                {
                    "httpMethod": "GET",
                    "path": "/cloudbed/scene/category/tree",
                    "name": "获取原子能力类目树"
                },
                {
                    "httpMethod": "POST",
                    "path": "/cloudbed/scene",
                    "name": "创建场景"
                },
                {
                    "httpMethod": "PUT",
                    "path": "/cloudbed/scene",
                    "name": "修改场景"
                },
                {
                    "httpMethod": "GET",
                    "path": "/cloudbed/scene",
                    "name": "场景分页"
                },
                {
                    "httpMethod": "DELETE",
                    "path": "/cloudbed/scene/*",
                    "name": "删除场景"
                },
                {
                    "httpMethod": "GET",
                    "path": "/cloudbed/scene/*",
                    "name": "获取场景"
                },
                {
                    "httpMethod": "POST",
                    "path": "/cloudbed/scene/duplication/*",
                    "name": "复制场景"
                },
                {
                    "httpMethod": "POST",
                    "path": "/cloudbed/scene/*/*",
                    "name": "添加场景标签"
                },
                {
                    "httpMethod": "DELETE",
                    "path": "/cloudbed/scene/*/*",
                    "name": "删除场景标签"
                },
                {
                    "httpMethod": "POST",
                    "path": "/cloudbed/scene/*",
                    "name": "运行场景"
                },
                {
                    "httpMethod": "POST",
                    "path": "/cloudbed/scene/verify/*",
                    "name": "运行场景前置校验"
                },
                {
                    "httpMethod": "PUT",
                    "path": "/cloudbed/scene/pipeline/sort",
                    "name": "阶段排序"
                },
                {
                    "httpMethod": "GET",
                    "path": "/cloudbed/experiment",
                    "name": "实验分页"
                },
                {
                    "httpMethod": "GET",
                    "path": "/cloudbed/experiment/detail/*",
                    "name": "实验详情"
                },
                {
                    "httpMethod": "DELETE",
                    "path": "/cloudbed/experiment/terminate/*",
                    "name": "终止实验"
                },
                {
                    "httpMethod": "POST",
                    "path": "/cloudbed/experiment/copy/*",
                    "name": "复制实验"
                },
                {
                    "httpMethod": "GET",
                    "path": "/cloudbed/experiment/report/*",
                    "name": "实验报告"
                },
                {
                    "httpMethod": "GET",
                    "path": "/cloudbed/experiment/job/events/*/*",
                    "name": "步骤日志"
                },
                {
                    "httpMethod": "GET",
                    "path": "/cloudbed/experiment/inspection/*/*",
                    "name": "查看状态检查记录"
                },
                {
                    "httpMethod": "GET",
                    "path": "/cloudbed/experiment/pre_execute/verify/*",
                    "name": "运行实验前置校验"
                },
                {
                    "httpMethod": "GET",
                    "path": "/cloudbed/experiment/re_execute/*",
                    "name": "运行实验"
                },
                {
                    "httpMethod": "GET",
                    "path": "/cloudbed/experiment/break/all",
                    "name": "一键停止所有实验"
                },
                {
                    "httpMethod": "GET",
                    "path": "/cloudbed/experiment/stop_loss/data",
                    "name": "获取止损指标监控数据"
                },
                {
                    "httpMethod": "POST",
                    "path": "/cloudbed/scene/pipeline/stage",
                    "name": "创建阶段"
                },
                {
                    "httpMethod": "PUT",
                    "path": "/cloudbed/scene/pipeline/stage",
                    "name": "修改阶段"
                },
                {
                    "httpMethod": "DELETE",
                    "path": "/cloudbed/scene/pipeline/stage/*",
                    "name": "删除阶段"
                },
                {
                    "httpMethod": "POST",
                    "path": "/cloudbed/scene/pipeline/job",
                    "name": "创建步骤"
                },
                {
                    "httpMethod": "PUT",
                    "path": "/cloudbed/scene/pipeline/job",
                    "name": "修改步骤"
                },
                {
                    "httpMethod": "DELETE",
                    "path": "/cloudbed/scene/pipeline/job/*",
                    "name": "删除步骤"
                },
                {
                    "httpMethod": "GET",
                    "path": "/cloudbed/scene/pipeline/job/*",
                    "name": "获取步骤"
                },
                {
                    "httpMethod": "GET",
                    "path": "/cloudbed/scene/pipeline/job/*/*",
                    "name": "获取步骤详情"
                },
                {
                    "httpMethod": "GET",
                    "path": "/cloudbed/scene/pipeline/jobs/*",
                    "name": "根据阶段获取步骤"
                },
                {
                    "httpMethod": "POST",
                    "path": "/cloudbed/scene/pipeline/inspection",
                    "name": "创建状态检查"
                },
                {
                    "httpMethod": "PUT",
                    "path": "/cloudbed/scene/pipeline/inspection",
                    "name": "修改状态检查"
                },
                {
                    "httpMethod": "GET",
                    "path": "/cloudbed/scene/pipeline/inspection/*",
                    "name": "获取状态检查"
                },
                {
                    "httpMethod": "DELETE",
                    "path": "/cloudbed/scene/pipeline/inspection/*",
                    "name": "删除状态检查"
                },
                {
                    "httpMethod": "POST",
                    "path": "/cloudbed/blade/ability",
                    "name": "批量添加原子能力"
                },
                {
                    "httpMethod": "GET",
                    "path": "/cloudbed/blade/ability/scope",
                    "name": "获取原子能力scope"
                },
                {
                    "httpMethod": "GET",
                    "path": "/cloudbed/blade/ability/filter",
                    "name": "过滤原子能力"
                },
                {
                    "httpMethod": "GET",
                    "path": "/cloudbed/blade/ability/template/*",
                    "name": "获取能力模板"
                },
                {
                    "httpMethod": "GET",
                    "path": "/cloudbed/blade/ability/template/filter",
                    "name": "过滤能力模板"
                },
                {
                    "httpMethod": "POST",
                    "path": "/cloudbed/rule/steadystate",
                    "name": "新建稳态规则"
                },
                {
                    "httpMethod": "PUT",
                    "path": "/cloudbed/rule/steadystate",
                    "name": "修改稳态规则"
                },
                {
                    "httpMethod": "GET",
                    "path": "/cloudbed/rule/steadystate",
                    "name": "稳态规则分页"
                },
                {
                    "httpMethod": "DELETE",
                    "path": "/cloudbed/rule/steadystate/*",
                    "name": "删除稳态规则"
                },
                {
                    "httpMethod": "GET",
                    "path": "/cloudbed/rule/steadystate/*",
                    "name": "获取稳态规则"
                },
                {
                    "httpMethod": "POST",
                    "path": "/cloudbed/ability/template",
                    "name": "增加能力模板"
                },
                {
                    "httpMethod": "PUT",
                    "path": "/cloudbed/ability/template",
                    "name": "修改能力模板"
                },
                {
                    "httpMethod": "GET",
                    "path": "/cloudbed/ability/template",
                    "name": "能力模板分页"
                },
                {
                    "httpMethod": "GET",
                    "path": "/cloudbed/ability/template/*",
                    "name": "获取能力模板"
                },
                {
                    "httpMethod": "DELETE",
                    "path": "/cloudbed/ability/template/*",
                    "name": "删除能力模板"
                },
                {
                    "httpMethod": "POST",
                    "path": "/cloudbed/simulation/script",
                    "name": "创建流量模拟(script)"
                },
                {
                    "httpMethod": "POST",
                    "path": "/cloudbed/simulation/har",
                    "name": "创建流量模拟(har)"
                },
                {
                    "httpMethod": "PUT",
                    "path": "/cloudbed/simulation/script",
                    "name": "修改流量模拟(script)"
                },
                {
                    "httpMethod": "PUT",
                    "path": "/cloudbed/simulation/har",
                    "name": "修改流量模拟(har)"
                },
                {
                    "httpMethod": "GET",
                    "path": "/cloudbed/simulation/script/*",
                    "name": "获取流量模拟(script)"
                },
                {
                    "httpMethod": "GET",
                    "path": "/cloudbed/simulation/har/*",
                    "name": "获取流量模拟(har)"
                },
                {
                    "httpMethod": "DELETE",
                    "path": "/cloudbed/simulation/*",
                    "name": "删除流量模拟"
                },
                {
                    "httpMethod": "GET",
                    "path": "/cloudbed/monitor/job",
                    "name": "获取步骤监控数据"
                }
            ],
            "menus": [
                {
                    "code": "CW-001",
                    "title": "概览",
                    "path": "/overview",
                    "icon": "storage",
                    "subMenus": []
                },
                {
                    "code": "CW-002",
                    "title": "场景管理",
                    "path": "/experiment",
                    "icon": "storage",
                    "subMenus": []
                },
                {
                    "code": "CW-003",
                    "title": "专家经验库",
                    "path": "/expert-library/",
                    "icon": "storage",
                    "subMenus": []
                },
                {
                    "code": "CW-004",
                    "title": "运维对抗",
                    "path": "/batch",
                    "icon": "storage",
                    "subMenus": []
                },
                {
                    "code": "CW-005",
                    "title": "定时计划",
                    "path": "/schedule",
                    "icon": "storage",
                    "subMenus": []
                },
                {
                    "code": "CW-006",
                    "title": "历史实验",
                    "path": "/experiment-history",
                    "icon": "storage",
                    "subMenus": []
                },
                {
                    "code": "CW-007",
                    "title": "爆破点管理",
                    "path": "/shot-point",
                    "icon": "storage",
                    "subMenus": []
                },
                {
                    "code": "CW-007",
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
                    "apiMatches": [],
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
	bytes, err := ioutil.ReadFile("/Users/farben/Desktop/upm.json")
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
