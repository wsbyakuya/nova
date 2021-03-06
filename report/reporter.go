package report

import (
	"fmt"
)

const (
	REPORT_FORMAT      = "测试总量: %d    pass: %d    fail: %d    测试通过率: %.2f%%\n"
	REPORT_TIME_FORMAT = "测试用时: %dms              平均用时: %dms\n"
)

type Messager struct {
	Pass       bool
	Url        string
	Body       string
	StatusCode int
	Time       int `单位：ms`
	ItemNum    int
}

type Reporter struct {
	size       int
	pass_count int
	time_count int
	timeout    int
	ReportText string
	Msgs       []*Messager
	IsSpread   bool
}

func NewReporter(size, timeout int) *Reporter {
	return &Reporter{
		size:       0,
		pass_count: 0,
		time_count: 0,
		timeout:    timeout,
		Msgs:       make([]*Messager, 0, size),
	}
}

func (r *Reporter) Add(m *Messager) {
	r.size++
	if m.Pass {
		r.pass_count++
	}
	r.time_count += m.Time

	r.Msgs = append(r.Msgs, m)
}

func (r *Reporter) Report() string {
	res := ""
	res = res + fmt.Sprintf(REPORT_FORMAT, r.size, r.pass_count, r.size-r.pass_count, float64(r.pass_count)/float64(r.size)*100)
	res = res + fmt.Sprintf(REPORT_TIME_FORMAT, r.time_count, r.time_count/r.size)
	return res
}
