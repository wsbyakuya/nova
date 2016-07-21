package main

import (
	"flag"
	"fmt"
	"os"
)

const version = "0.5.0"

const (
	CMD_NO_CONFIG = iota
	CMD_API_CONFIG
	CMD_ENV_CONFIG
	CMD_BOTH_CONFIG
)

var commands = []*Command{}
var (
	GlobalPath         string
	argExport, argOpen bool
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
	if requestNum&2 == 2 {
		loadConfig("env.cfg")
	}
	if requestNum&1 == 1 {
		loadParas("api.cfg")
	}
}

//程序交互通用函数
func FailAndExit(err error) {
	fmt.Println(err)
	os.Exit(1)
}

func AskAndScan(q string) string {
	var str string
	fmt.Print(q)
	fmt.Scanln(&str)
	return str
}

//检测字符串切片是否包含某字符串
type Args []string

func (a Args) Contains(str string) bool {
	for _, v := range a {
		if v == str {
			return true
		}
	}
	return false
}

func (a Args) Parse() map[byte]bool {
	m := make(map[byte]bool, 5)
	fmt.Println(a)
	for _, v := range a {
		ss := []byte(v)
		fmt.Println(ss)
		if len(ss) > 1 && ss[0] == '-' {
			for _, k := range ss[1:] {
				fmt.Println("map key ", k)
				m[k] = true
			}
		}
	}
	return m
}
