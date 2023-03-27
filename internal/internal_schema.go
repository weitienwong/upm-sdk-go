package internal

// Author weitien

type User struct {
	UserId   int64    `json:"userId"`
	Username string   `json:"username"`
	Nickname string   `json:"nickname"`
	Apps     []*App   `json:"apps,omitempty"`
	Groups   []*Group `json:"groups,omitempty"`
	Panels   []*Panel `json:"panels,omitempty"`
}

type Menu struct {
	Code     string  `json:"code,omitempty"`
	Title    string  `json:"title,omitempty"`
	Path     string  `json:"path,omitempty"`
	IconPath string  `json:"iconPath,omitempty"`
	SubMenus []*Menu `json:"subMenus,omitempty"`
}

type App struct {
	Code        string  `json:"code"`
	Name        string  `json:"name"`
	Url         string  `json:"url"`
	IconPath    string  `json:"iconPath,omitempty"`
	Description string  `json:"description,omitempty"`
	Menus       []*Menu `json:"menus,omitempty"`
	Apis        []*Api  `json:"apis"`
}

type Group struct {
	GroupId     int64    `json:"groupId"`
	Name        string   `json:"name"`
	IconPath    string   `json:"iconPath"`
	Description string   `json:"description,omitempty"`
	AppCodes    []string `json:"appCodes"`
}

type Panel struct {
	Name     string  `json:"name"`
	GroupIds []int64 `json:"groupIds"`
}

type Api struct {
	RequestMethod string `json:"reqMethod"`
	Url           string `json:"url"`
}
