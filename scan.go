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

	if argsMap['c'] {
		Cookies = loadFile("cookies.cfg")
	}
	if argsMap['h'] {
		Header = loadFile("header.cfg")
	}

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
	if argsMap['e'] {
		reporter.ExportHTML(argsMap['o'])
	}
}

func testItem(url string, r *report.Reporter) {
	req, err := fetch.NewNovaRequest("GET", url)
	if err != nil {
		FailAndExit(err)
	}
	req.SetCookies(Cookies)
	req.SetHeader(Header)

	res, err2 := req.Do()
	if err2 != nil {
		FailAndExit(err)
	}
	itemNum := handle.ItemsNum(res.Body)
	pass := false
	if itemNum > 0 {
		pass = true
	}
	m := report.Messager{
		Pass:       pass,
		Url:        res.Url,
		Body:       string(res.Body),
		Status:     res.Status,
		StatusCode: res.StatusCode,
		Time:       res.Time,
		ItemNum:    itemNum,
	}
	r.Add(&m)
}
