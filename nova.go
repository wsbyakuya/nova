package main

import (
	"flag"
	"fmt"
	"os"
)

const version = "0.2.2"

var commands = []*Command{}
var GlobalPath string

type Command struct {
	Run  func([]string)
	Name string
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
			cmd.Run(args[1:])
			return
		}
	}
}

func FailAndExit(err error) {
	fmt.Println(err)
	os.Exit(1)
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
