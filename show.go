package main

import (
	"fmt"
)

var (
	cmdVersion = &Command{
		Name: "version",
	}
	cmdApi = &Command{
		Name: "api",
	}
	cmdEnv = &Command{
		Name: "env",
	}
	cmdHelp = &Command{
		Name: "help",
	}
)

func init() {
	commands = append(commands, cmdVersion, cmdApi, cmdEnv, cmdHelp)
	cmdVersion.Run = showVersion
	cmdApi.Run = showApi
	cmdEnv.Run = showEnv
	cmdHelp.Run = showUsage
}

func showVersion(args []string) {
	fmt.Println("nova version " + version)
}

const PARA_FORMAT = "para: %-20svalues:%v\n"

func showApi(args []string) {
	fmt.Println("api:  " + Api)
	for k, v := range Paras {
		fmt.Printf(PARA_FORMAT, k, v)
	}
}

const ENV_FORMAT = "%10s%v\n"

func showEnv(args []string) {
	fmt.Printf(ENV_FORMAT, "Host1:", Host1)
	fmt.Printf(ENV_FORMAT, "Host2:", Host2)
	fmt.Printf(ENV_FORMAT, "Hosts:", Hosts)
	fmt.Printf(ENV_FORMAT, "Port:", Port)
	fmt.Printf(ENV_FORMAT, "MaxDiff:", MaxDiff)
	fmt.Printf(ENV_FORMAT, "Timeout:", Timeout)
}

func showUsage(args []string) {
	msg := `命令参数
nova version 显示当前版本号
nova env     显示测试环境设置
nova api     显示待测api参数信息`
	fmt.Println(msg)
}
