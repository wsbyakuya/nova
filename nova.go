package main

import (
	"flag"
	"fmt"
	"os"
)

const version = "0.6.0"

const (
	CMD_NO_CONFIG = iota
	CMD_API_CONFIG
	CMD_ENV_CONFIG
	CMD_BOTH_CONFIG
)

var commands = []*Command{}
var (
	GlobalPath string
)

type Command struct {
	Run           func([]string)
	Name          string
	ConfigRequest int
}

func init() {
	GlobalPath = ""
}

func main() {
	flag.Parse()
	args := flag.Args()
	if len(args) == 0 {
		os.Exit(2)
	}
	for _, cmd := range commands {
		if cmd.Name == args[0] {
			loadConfigFiles(cmd.ConfigRequest)
			cmd.Run(args[1:])
			return
		}
	}
}

func loadConfigFiles(requestNum int) {
	if requestNum&CMD_ENV_CONFIG == CMD_ENV_CONFIG {
		loadConfig("env.cfg")
	}
	if requestNum&CMD_API_CONFIG == CMD_API_CONFIG {
		loadParas("api.cfg")
	}
}

//程序交互通用函数
func FailAndExit(err error) {
	fmt.Println(err)
	os.Exit(1)
}

func AskAndScan(question string) string {
	var str string
	fmt.Print(question)
	fmt.Scanln(&str)
	return str
}

//检测字符串切片是否包含某字符串
type Args []string

func (a Args) Contains(strs ...string) bool {
	for _, v := range a {
		for _, k := range strs {
			if v == k {
				return true
			}
		}
	}
	return false
}

func (a Args) Parse() map[byte]bool {
	m := make(map[byte]bool, 5)
	for _, v := range a {
		ss := []byte(v)
		if len(ss) > 1 && ss[0] == '-' {
			for _, k := range ss[1:] {
				m[k] = true
			}
		}
	}
	return m
}
