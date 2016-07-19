package main

import (
	"fmt"
	"github.com/wsbyakuya/nova/fetch"
	"github.com/wsbyakuya/nova/handle"
	"github.com/wsbyakuya/nova/report"
	"strings"
)

const (
	PASS_REPORT_FORMAT = "测试总量: %d    pass: %d    fail: %d    测试通过率: %%%.3f \n"
	PASS_TIME_FORMAT   = "测试用时: %.3fs    平均用时: %.3fs\n"
)

var cmdCompare = &Command{
	Name:          "compare",
	ConfigRequest: CMD_BOTH_CONFIG,
}

func init() {
	commands = append(commands, cmdCompare)
	cmdCompare.Run = compareTestAll
}

func compareTestAll(args []string) {
	//开始测试
	var uri1, uri2 string
	if len(args) > 0 {
		//根据参数选择host
	}
	//组装url
	//默认对比host1和host2
	uri1 = "http://" + Host1
	uri2 = "http://" + Host2
	if Port != "" {
		uri1 = uri1 + ":" + Port
		uri2 = uri2 + ":" + Port
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
	fmt.Printf("开始测试 %s  vs  %s\n%s\n\n", uri1, uri2, Api)
	if testSize > 0 {
		for i, v := range fullList {
			fmt.Printf("\r正在测试    %d/%d", i+1, testSize)
			url1 := uri1 + const_api + v
			url2 := uri2 + const_api + v
			compareTestItem(url1, url2, reporter)
		}
		fmt.Println("\n测试完成")
	} else {
		url1 := uri1 + const_api
		url2 := uri2 + const_api
		compareTestItem(url1, url2, reporter)
		fmt.Println("测试完成")
	}
	fmt.Println(reporter.Report())
}

func compareTestItem(url1, url2 string, r *report.Reporter) {
	res1, err1 := fetch.HttpGet(url1)
	if err1 != nil {
		FailAndExit(err1)
	}
	res2, err2 := fetch.HttpGet(url2)
	if err2 != nil {
		FailAndExit(err2)
	}

	m := report.Messager{
		Pass:       handle.CompareBody(res1.Body, res2.Body, MaxDiff),
		Url:        url2,
		Time:       res2.Time,
		StatusCode: res2.StatusCode,
	}
	r.Add(&m)
}
