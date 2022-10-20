package RSL_Setting

import (
	"RedStoneLauncher/RSL_Log"
	"encoding/json"
	"io/ioutil"
	"os"
)

func InitSettings() {
	if _, iSErrA := os.Stat(ConfigPath); iSErrA != nil {
		if os.IsNotExist(iSErrA) {
			fileContent, err := json.Marshal(DefaultConfig)
			if err = ioutil.WriteFile(ConfigPath, fileContent, 0666); err != nil {
				RSL_Log.LogWarning(err.Error())
			} else {
				RSL_Log.LogInfo("Successfully init settings!")
			}
			CurrentConfig = DefaultConfig
			return
		}
	}
	//读取文件
	fileContent, err := ioutil.ReadFile(ConfigPath)
	if err != nil {
		RSL_Log.LogWarning(err.Error())
		return
	}
	if err := json.Unmarshal(fileContent, &CurrentConfig); err != nil {
		RSL_Log.LogWarning(err.Error())
	} else {
		RSL_Log.LogInfo("Successfully init settings!")
	}
}
