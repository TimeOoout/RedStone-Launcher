package RSL_Setting

var CurrentConfig = DefaultConfig
var DefaultConfig = LauncherConfig{
	/*用户数据*/
	UserInfo: UserInfo{
		//角色名
		DisplayName: "",
		//账户（邮箱之类的）
		Account: "",
		//用户类型（Mojang|微软/红石/离线）
		Mod:            "",
		Uuid:           "",
		UserProperties: nil,
		AccessToken:    "",
		ProfileProperties: profileProperties{
			Textures:           "",
			UploadableTextures: "",
		},
		//用户是否被选中
		Selected: false,
	},
	/*项目设置*/
	LauncherSettings: LauncherSettings{
		//可用Java列表
		JavaList: nil,
		//所有游戏路径
		GameDir: nil,
		//用户设置
		Player: nil,
		//游戏设置
		Config: config{
			//全局设置
			Global: globalConfig{
				//是否使用全局设置
				UseGlobal: true,
				//最大内存
				MaxMemory: 1024,
				//自动分配内存
				AutoMemory: true,
				//Java参数
				JavaArgs: "",
				//Java版本
				Java: "",
				//本地Java路径
				DefaultJavaPath: "",
			},
			//自定义配置
			Private: nil,
		},
		Download: downloadSettings{
			//下载重试次数
			RetryTimes: 3,
			//单个文件下载线程数
			SingleThreads: 4,
			//单次瞎子啊任务线程数
			DownloadThreads: 100,
			//下载源
			Source: downloadSource{
				Official: sourceURL{
					VersionList_v1: "http://launchermeta.mojang.com/mc/game/version_manifest.json",
					VersionList_v2: "http://launchermeta.mojang.com/mc/game/version_manifest_v2.json",
					Index:          "https://launchermeta.mojang.com/",
					Launcher:       "https://launcher.mojang.com/",
					Assets:         "http://resources.download.minecraft.net/",
					Libraries:      "https://libraries.minecraft.net/",
				},
				BMCL: sourceURL{
					VersionList_v1: "https://bmclapi2.bangbang93.com/mc/game/version_manifest.json",
					VersionList_v2: "https://bmclapi2.bangbang93.com/mc/game/version_manifest_v2.json",
					Index:          "https://bmclapi2.bangbang93.com/",
					Launcher:       "https://bmclapi2.bangbang93.com/",
					Assets:         "https://bmclapi2.bangbang93.com/assets/",
					Libraries:      "https://bmclapi2.bangbang93.com/maven/",
				},
				MCBBS: sourceURL{
					VersionList_v1: "https://download.mcbbs.net/mc/game/version_manifest.json",
					VersionList_v2: "https://download.mcbbs.net/mc/game/version_manifest_v2.json",
					Index:          "https://download.mcbbs.net/",
					Launcher:       "https://download.mcbbs.net/",
					Assets:         "https://download.mcbbs.net/assets/",
					Libraries:      "https://download.mcbbs.net/maven/",
				},
				//预留了一个自定义源
				Others: sourceURL{
					VersionList_v1: "",
					VersionList_v2: "",
					Index:          "",
					Launcher:       "",
					Assets:         "",
					Libraries:      "",
				},
			},
			//下载模式，0为官方，1为BMCL，2为MCBBS（推荐），3为自定义源
			Mod: 2,
		},
		AutoLog: true,
		About: aboutInfo{
			Version: "0.1.1_Alpha",
			Url:     "https://mcskin.cn/",
		},
	},
}
