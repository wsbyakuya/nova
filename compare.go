package main

import (
	"fmt"
	"github.com/wsbyakuya/nova/fetch"
	"github.com/wsbyakuya/nova/handle"
	"strings"
	"time"
)

const (
	PASS_REPORT_FORMAT = "测试总量: %d    pass: %d    fail: %d    测试通过率: %%%.3f \n"
	PASS_TIME_FORMAT   = "测试用时: %.3fs    平均用时: %.3fs\n"
)

var cmdCompare = &Command{
	Name: "compare",
}

func init() {
	commands = append(commands, cmdCompare)
	cmdCompare.Run = compareTestAll
}

func compareTestAll(args []string) {
	//开始测试
	begin_time := time.Now()
	var uri1, uri2 string
	if len(args) > 0 {
		//根据参数选择host
	}
	sum, pass_count := 0, 0
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
		const_api = const_api + strings.Join(ConstParas, "&") + "&"
	}

	fullList := fetch.GetFullList(Paras)

	for _, v := range fullList {
		interface_test := const_api + v
		if pass := compareTestItem(uri1, uri2, interface_test); pass {
			pass_count++
		}
		sum++
	}

	end_time := time.Now()
	fmt.Printf(PASS_REPORT_FORMAT, sum, pass_count, sum-pass_count, float64(pass_count)/float64(sum)*100)
	time_use := end_time.Sub(begin_time).Seconds()
	fmt.Printf(PASS_TIME_FORMAT, time_use, time_use/float64(sum))
}

func compareTestItem(uri1, uri2, ifs string) bool {
	fmt.Println("正在测试接口: " + ifs)
	url1 := uri1 + ifs
	url2 := uri2 + ifs

	res1, err1 := fetch.HttpGet(url1)
	if err1 != nil {
		FailAndExit(err1)
	}
	res2, err2 := fetch.HttpGet(url2)
	if err2 != nil {
		FailAndExit(err2)
	}

	if pass, message := handle.CompareBody(res1.Body, res2.Body, MaxDiff); pass {
		fmt.Println("Pass")
		return true
	} else {
		fmt.Println("Failed: " + message)
		return false
	}
}
