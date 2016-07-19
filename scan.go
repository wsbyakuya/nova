package main

import (
	"fmt"
	"github.com/wsbyakuya/nova/fetch"
	"github.com/wsbyakuya/nova/handle"
	"github.com/wsbyakuya/nova/report"
	"strings"
)

var cmdScan = &Command{
	Name: "scan",
}

func init() {
	commands = append(commands, cmdScan)
	cmdScan.Run = scan
}

func scan(args []string) {
	uri := "http://" + Host1
	if Port != "" {
		uri = uri + ":" + Port
	}
	const_api := Api
	if len(Paras) > 0 || len(ConstParas) > 0 {
		const_api = Api + "?"
	}
	if len(ConstParas) > 0 {
		const_api = const_api + strings.Join(ConstParas, "&")
		if len(Paras) > 0 {
			const_api += "&"
		}
	}

	fullList := fetch.GetFullList(Paras)
	testSize := len(fullList)

	reporter := report.NewReporter(testSize, Timeout)

	fmt.Printf("开始测试 %s\n%s\n\n", uri, Api)
	if testSize > 0 {
		for i, p := range fullList {
			testItem(uri+const_api+p, reporter)
			fmt.Printf("\r正在测试    %d/%d", i+1, testSize)
		}
		fmt.Println("\n测试完成")
	} else {
		testItem(uri+const_api, reporter)
		fmt.Println("测试完成")
	}
	fmt.Println(reporter.Report())
	reporter.ExportHTML()
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
