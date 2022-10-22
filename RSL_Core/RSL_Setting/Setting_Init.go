package RSL_Setting

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

func InitSettings() error {
	if _, iSErrA := os.Stat(ConfigPath); iSErrA != nil {
		if os.IsNotExist(iSErrA) {
			fileContent, err := json.Marshal(DefaultConfig)
			if err = ioutil.WriteFile(ConfigPath, fileContent, 0666); err != nil {
				logWarningInitSettings(err.Error())
				return err
			} else {
				logInfoInitSetting()
			}
			CurrentConfig = DefaultConfig
			return nil
		}
	}
	//读取文件
	fileContent, err := ioutil.ReadFile(ConfigPath)
	if err != nil {
		logWarningInitSettings(err.Error())
		return err
	}
	if err := json.Unmarshal(fileContent, &CurrentConfig); err != nil {
		logWarningInitSettings(err.Error())
		return err
	} else {
		logInfoInitSetting()
	}
	return nil
}
