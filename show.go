package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

var (
	cmdVersion = &Command{
		Name:          "version",
		ConfigRequest: CMD_NO_CONFIG,
	}
	cmdApi = &Command{
		Name:          "api",
		ConfigRequest: CMD_API_CONFIG,
	}
	cmdEnv = &Command{
		Name:          "env",
		ConfigRequest: CMD_ENV_CONFIG,
	}
	cmdHelp = &Command{
		Name:          "help",
		ConfigRequest: CMD_NO_CONFIG,
	}
	cmdShow = &Command{
		Name:          "show",
		ConfigRequest: CMD_BOTH_CONFIG,
	}
	cmdEdit = &Command{
		Name:          "edit",
		ConfigRequest: CMD_API_CONFIG,
	}
)

func init() {
	commands = append(commands, cmdVersion, cmdApi, cmdEnv, cmdHelp, cmdShow, cmdEdit)
	cmdVersion.Run = showVersion
	cmdApi.Run = showApi
	cmdEnv.Run = showEnv
	cmdHelp.Run = showUsage
	cmdShow.Run = showAll
	cmdEdit.Run = edit
}

func showVersion(args []string) {
	fmt.Println("nova version " + version)
}

func showAll(args []string) {
	fmt.Println("============environment============")
	showEnv(args)
	fmt.Println("================api================")
	showApi(args)
}

const PARA_FORMAT = "%s = %v\n"

func showApi(args []string) {
	fmt.Println("api:  " + Api)
	for _, v := range ConstParas {
		ss := strings.Split(v, "=")
		if len(ss) >= 2 {
			fmt.Printf(PARA_FORMAT, ss[0], ss[1])
		}
	}
	for k, v := range Paras {
		str := strings.Join(v, "|")
		fmt.Printf(PARA_FORMAT, k, str)
	}

	//show cookies
	if len(Cookies) > 0 {
		fmt.Println("Cookies:")
		for k, v := range Cookies {
			fmt.Printf(PARA_FORMAT, k, v)
		}
	}
	//show header
	if len(Header) > 0 {
		fmt.Println("Header:")
		for k, v := range Header {
			fmt.Printf(PARA_FORMAT, k, v)
		}
	}
}

const ENV_FORMAT = "%s = %v\n"

func showEnv(args []string) {
	fmt.Printf(ENV_FORMAT, "Host1", Host1)
	fmt.Printf(ENV_FORMAT, "Host2", Host2)
	fmt.Printf(ENV_FORMAT, "Port", Port)
	fmt.Printf(ENV_FORMAT, "MaxDiff", MaxDiff)
	fmt.Printf(ENV_FORMAT, "Timeout", Timeout)
}

func showUsage(args []string) {
	msg := `nova is a tool for RESTful api test.

Uasge:

	nova command [arguments]

Commands:

	nova api         显示待测api参数信息
	nova compare     对比Host1和Host2同一接口返回量
	nova env         显示测试环境设置
	nova scan        扫描接口的所有参量组合
	nova version     显示当前版本号`
	fmt.Println(msg)
}

func edit(args []string) {
	cmd := exec.Command("C:/Program Files/Sublime Text 3/sublime_text.exe")
	file := "api.cfg"
	path, err := os.Getwd()
	if err == nil {
		file = path + "/" + file
	}
	cmd.Args = append(cmd.Args, file)
	cmd.Start()
}
