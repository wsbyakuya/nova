package main

import (
	"fmt"
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
)

func init() {
	commands = append(commands, cmdVersion, cmdApi, cmdEnv, cmdHelp, cmdShow)
	cmdVersion.Run = showVersion
	cmdApi.Run = showApi
	cmdEnv.Run = showEnv
	cmdHelp.Run = showUsage
	cmdShow.Run = showAll
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
		str := strings.Join(v, ",")
		fmt.Printf(PARA_FORMAT, k, str)
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
