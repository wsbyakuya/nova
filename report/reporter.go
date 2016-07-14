package report

import (
	"fmt"
)

const (
	REPORT_FORMAT      = "测试总量: %d    pass: %d    fail: %d    测试通过率: %%%.3f\n"
	REPORT_TIME_FORMAT = "测试用时: %dms              平均用时: %dms\n"
)

type Messager struct {
	Pass       bool
	Url        string
	Body       []byte
	StatusCode int
	Time       int `单位：ms`
}

type Reporter struct {
	size       int
	pass_count int
	time_count int
	timeout    int
	msgs       []Messager
}

func NewReporter(size, timeout int) *Reporter {
	return &Reporter{
		size:       size,
		pass_count: 0,
		time_count: 0,
		timeout:    timeout,
		msgs:       make([]Messager, 0, size),
	}
}

func (r *Reporter) Add(m Messager) {
	if m.Pass {
		r.pass_count++
	}
	r.time_count += m.Time

	r.msgs = append(r.msgs, m)
}

func (r *Reporter) Report() string {
	res := ""
	res = res + fmt.Sprintf(REPORT_FORMAT, r.size, r.pass_count, r.size-r.pass_count, float64(r.pass_count)/float64(r.size)*100)
	res = res + fmt.Sprintf(REPORT_TIME_FORMAT, r.time_count, r.time_count/r.size)
	return res
}
