package main

import (
	// "fmt"
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
	newApi := args[0]
	if []byte(newApi)[0] != '/' {
		newApi = "/" + newApi
	}

	dirName := []byte(strings.Replace(newApi, "/", "_", -1))
	if len(dirName) > 0 && dirName[0] == '_' {
		dirName = dirName[1:]
	}
	os.Mkdir(string(dirName), os.ModeDir)

}
