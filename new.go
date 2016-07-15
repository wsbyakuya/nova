package main

import (
	"fmt"
	"os"
	"strings"
)

var cmdNew = &Command{
	Name: "new",
}

func init() {
	commands = append(commands, cmdNew)
	cmdNew.Run = newFunc
}

func newFunc(args []string) {
	var newApi string
	if len(args) > 0 {
		newApi = args[0]
	} else {
		q := "请输入要创建的api --> "
		newApi = AskAndScan(q)
	}

	if []byte(newApi)[0] != '/' {
		newApi = "/" + newApi
	}

	dirName := []byte(strings.Replace(newApi, "/", "_", -1))
	if len(dirName) > 0 && dirName[0] == '_' {
		dirName = dirName[1:]
	}
	dirStringName := string(dirName)
	var err error
	err = os.Mkdir(dirStringName, os.ModeDir)
	if err != nil {
		if os.IsExist(err) {
			q := "该测试目录已存在，是否覆盖测试文件(y/n)？"
			if AskAndScan(q) != "y" {
				return
			}
		} else {
			FailAndExit(err)
		}
	}

	err = newApiConfigFile(dirStringName+"\\api.cfg", newApi)
	if err != nil {
		FailAndExit(err)
	}
}

func newApiConfigFile(file, api string) error {
	f, err := os.Create(file)
	if err != nil {
		return err
	}
	defer f.Close()

	context := `# API配置

api = %s

# 参数配置`

	context = fmt.Sprintf(context, api)
	_, err2 := f.Write([]byte(context))
	if err2 != nil {
		return err2
	}
	return nil
}
