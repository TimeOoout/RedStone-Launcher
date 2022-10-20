package RSL_Setting

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

func InitSettings() {
	readSettings()
}

func readSettings() {
	fileContent, err := json.Marshal(DefaultConfig)
	println(fileContent)
	if err = ioutil.WriteFile(ConfigPath, fileContent, 0666); err != nil {
		fmt.Println("Writefile Error =", err)
		return
	}
	//读取文件
	fileContent, err = ioutil.ReadFile(ConfigPath)
	if err != nil {
		fmt.Println("Read file err =", err)
		return
	}
	if err := json.Unmarshal(fileContent, &CurrentConfig); err != nil {
		fmt.Println("Read file error =", err)
	} else {
		fmt.Println("Read file success =", CurrentConfig)
	}
}
