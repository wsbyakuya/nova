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
