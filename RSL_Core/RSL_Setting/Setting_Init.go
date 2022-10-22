package RSL_Setting

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

func InitSettings() {
	if _, iSErrA := os.Stat(ConfigPath); iSErrA != nil {
		if os.IsNotExist(iSErrA) {
			fileContent, err := json.Marshal(DefaultConfig)
			if err = ioutil.WriteFile(ConfigPath, fileContent, 0666); err != nil {
				logWarningInitSettings(err.Error())
			} else {
				logInfoInitSetting()
			}
			CurrentConfig = DefaultConfig
			return
		}
	}
	//读取文件
	fileContent, err := ioutil.ReadFile(ConfigPath)
	if err != nil {
		logWarningInitSettings(err.Error())
		return
	}
	if err := json.Unmarshal(fileContent, &CurrentConfig); err != nil {
		logWarningInitSettings(err.Error())
	} else {
		logInfoInitSetting()
	}
}
