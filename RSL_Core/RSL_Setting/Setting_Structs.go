package RSL_Setting

/*Launcher类*/

type LauncherConfig struct {
	UserInfo         UserInfo
	LauncherSettings LauncherSettings
}

/*用户信息*/

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

type profileProperties struct {
	Textures           string `json:"textures"`
	UploadableTextures string `json:"uploadableTextures"`
}

/*启动器设置*/
type LauncherSettings struct {
	JavaList interface{}      `json:"javaList"`
	GameDir  interface{}      `json:"gameDir"`
	Player   interface{}      `json:"player"`
	Config   config           `json:"config"`
	Download downloadSettings `json:"download"`
	AutoLog  bool             `json:"autoLog"`
	About    aboutInfo        `json:"about"`
}

type gameDir struct {
	Path        string `json:"path"`
	DisplayName string `json:"displayName"`
}

type playerInfo struct {
}

type config struct {
	Global  globalConfig `json:"global"`
	Private interface{}  `json:"private"`
}

type globalConfig struct {
	UseGlobal       bool   `json:"useGlobal"`
	MaxMemory       int    `json:"maxMemory"`
	AutoMemory      bool   `json:"autoMemory"`
	JavaArgs        string `json:"javaArgs"`
	Java            string `json:"java"`
	DefaultJavaPath string `json:"defaultJavaPath"`
}

type downloadSettings struct {
	RetryTimes      int            `json:"retryTimes"`
	SingleThreads   int            `json:"singleThreads"`
	DownloadThreads int            `json:"downloadThreads"`
	Source          downloadSource `json:"source"`
	Mod             int            `json:"mod"`
}
type downloadSource struct {
	Official sourceURL `json:"official"`
	BMCL     sourceURL `json:"BMCL"`
	MCBBS    sourceURL `json:"MCBBS"`
	Others   sourceURL `json:"Others"`
}

type sourceURL struct {
	VersionList_v1 string `json:"VersionList_V1"`
	VersionList_v2 string `json:"version_list___v_2"`
	Index          string `json:"index"`
	Launcher       string `json:"launcher"`
	Assets         string `json:"assets"`
	Libraries      string `json:"libraries"`
}

type aboutInfo struct {
	Version string `json:"version"`
	Url     string `json:"url"`
}
