package RSL_Player

import (
	"io/ioutil"
	"os"
)

func SetUsers() {

}

func GetUsers() error {
	//判断是否存在该路径，若不存在则自动创建
	if _, iSErrA := os.Stat(ConfigFolder); iSErrA != nil {
		if os.IsNotExist(iSErrA) {
			os.MkdirAll(ConfigFolder, 7777)
			return nil
		}
	} else {
		rd, err := ioutil.ReadDir(ConfigFolder)
		if err != nil {
			logWarningReadDir(err)
			return err
		}
		///未完成

		var s []string
		for _, fi := range rd {
			if !fi.IsDir() {
				fullName := ConfigFolder + fi.Name()
				s = append(s, fullName)
			}
		}

		///
		return nil
	}
	return nil
}
