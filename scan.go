package main

import (
	"fmt"
	"strings"

	"github.com/wsbyakuya/nova/fetch"
	"github.com/wsbyakuya/nova/handle"
	"github.com/wsbyakuya/nova/report"
)

var cmdScan = &Command{
	Name:          "scan",
	ConfigRequest: CMD_BOTH_CONFIG,
}

func init() {
	commands = append(commands, cmdScan)
	cmdScan.Run = scan
}

func scan(args []string) {
	argsMap := Args(args).Parse()

	uri := "http://" + Host1
	if Port != "" {
		uri = uri + ":" + Port
	}
	constAPI := Api
	if len(Paras) > 0 || len(ConstParas) > 0 {
		constAPI = Api + "?"
	}
	if len(ConstParas) > 0 {
		constAPI = constAPI + strings.Join(ConstParas, "&")
		if len(Paras) > 0 {
			constAPI += "&"
		}
	}

	fullList := fetch.GetFullList(Paras)
	testSize := len(fullList)

	reporter := report.NewReporter(testSize, Timeout)

	fmt.Printf("\n开始测试 %s\n%s\n\n", uri, Api)
	if testSize > 0 {
		for i, p := range fullList {
			testItem(uri+constAPI+p, reporter)
			fmt.Printf("\r正在测试    %d/%d", i+1, testSize)
		}
		fmt.Println("\n测试完成")
	} else {
		testItem(uri+constAPI, reporter)
		fmt.Println("测试完成")
	}
	fmt.Println(reporter.Report())

	if argsMap['b'] {
		reporter.IsSpread = true
	}
	if argsMap['h'] || argsMap['e'] {
		reporter.ExportHTML(argsMap['o'])
	}
}

func testItem(url string, r *report.Reporter) {
	res, err := fetch.HttpGet(url)
	if err != nil {
		panic(err)
	}
	m := report.Messager{
		Pass:       !handle.IsEmpty(res.Body),
		Url:        res.Url,
		Body:       string(res.Body),
		StatusCode: res.StatusCode,
		Time:       res.Time,
	}
	r.Add(&m)
}
