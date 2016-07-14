package fetch

import (
	"errors"
	"io/ioutil"
	"net/http"
	"time"
)

type Response struct {
	Body       []byte
	Url        string
	StatusCode int
	Time       int `ms`
}

func HttpGet(urlStr string) (response *Response, err error) {
	req, err := http.NewRequest("GET", urlStr, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Add("cache-control", "no-cache")
	start := time.Now()

	res, err2 := http.DefaultClient.Do(req)
	if err2 != nil {
		return nil, errors.New("无法获取连接：" + urlStr + "\n请检查服务器运行情况\n")
	}
	time_use := time.Since(start).Nanoseconds() / 1E6

	defer res.Body.Close()

	body, err3 := ioutil.ReadAll(res.Body)
	if err3 != nil {
		return nil, err3
	}
	r := Response{
		Body:       body,
		Url:        urlStr,
		StatusCode: res.StatusCode,
		Time:       int(time_use),
	}
	return &r, nil
}
