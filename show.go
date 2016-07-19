package main

import (
	"fmt"
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
nova api     显示待测api参数信息
nova compare 对比Host1和Host2同一接口返回量
nova env     显示测试环境设置
nova scan    扫描接口的所有参量组合
nova version 显示当前版本号
`
	fmt.Println(msg)
}
