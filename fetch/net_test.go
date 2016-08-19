package fetch

import (
	"testing"
)

func TestHttpGet(t *testing.T) {
	res, err := HttpGet("http://m.api.zhe800.com/deals/dailyten/v1")
	if err != nil {
		t.Error(err)
	}
	t.Log(res.StatusCode)
	t.Log(res.Status)
	t.Log(res.Time)
}

func TestHttpPost(t *testing.T) {
	res, err := HttpPost("http://localhost:8080/post", "hello world")
	if err != nil {
		t.Error(err)
	}
	t.Log(res.StatusCode)
	t.Log(res.Status)
	t.Log(res.Time)
}
