package main

import (
	"fmt"
	"strings"
)

var cmdNew = &Command{
	Name: "new",
}

func init() {
	commands = append(commands, cmdScan)
	cmdScan.Run = newCmd
}

func newCmd(args []string) {

}
