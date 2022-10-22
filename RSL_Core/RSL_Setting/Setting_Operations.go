package RSL_Setting

import (
	"RedStoneLauncher/RSL_Core/RSL_Log"
	"encoding/json"
	"io/ioutil"
)

func GetSettings() {
	fileContent, err := ioutil.ReadFile(ConfigPath)
	if err != nil {
		RSL_Log.LogWarning(err.Error())
		return
	}
	if err := json.Unmarshal(fileContent, &CurrentConfig); err != nil {
		RSL_Log.LogWarning(err.Error())
	} else {
		RSL_Log.LogInfo("Successfully Get settings!")
	}
}

func SetSettings() {
	fileContent, err := json.Marshal(DefaultConfig)
	if err = ioutil.WriteFile(ConfigPath, fileContent, 0666); err != nil {
		RSL_Log.LogWarning(err.Error())
	} else {
		RSL_Log.LogInfo("Successfully Set settings!")
	}
}
