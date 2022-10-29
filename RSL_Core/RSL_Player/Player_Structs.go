package RSL_Player

// Players  所有用户
type Users struct {
	Info interface{}
}

// UserInfo 用户信息
type UserInfo struct {
	DisplayName       string            `json:"display_name"`
	Account           string            `json:"account"`
	Mod               string            `json:"mod"`
	Uuid              string            `json:"uuid"`
	UserProperties    interface{}       `json:"userProperties"`
	AccessToken       string            `json:"accessToken"`
	ProfileProperties profileProperties `json:"profileProperties"`
	Selected          bool              `json:"selected"`
}

// profileProperties 用户皮肤数据信息
type profileProperties struct {
	Textures           string `json:"textures"`
	UploadableTextures string `json:"uploadableTextures"`
}
