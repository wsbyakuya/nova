package fetch

import (
	"errors"
	"io/ioutil"
	"net/http"
	"time"
)

type NovaRequest struct {
	request *http.Request
}

func NewNovaRequest(method, urlStr string) (*NovaRequest, error) {
	req, err := http.NewRequest(method, urlStr, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Add("Cache-Control", "no-cache")
	return &NovaRequest{
		request: req,
	}, nil
}

func (r *NovaRequest) AddCookie(name, value string) {
	r.request.AddCookie(&http.Cookie{
		Name:  name,
		Value: value,
	})
}

func (r *NovaRequest) SetCookies(cookies map[string]string) {
	if len(cookies) == 0 {
		return
	}
	for k, v := range cookies {
		r.AddCookie(k, v)
	}
}

func (r *NovaRequest) Do() (response *Response, err error) {
	start := time.Now()

	res, err := http.DefaultClient.Do(r.request)
	if err != nil {
		return nil, errors.New("无法获取连接：" + r.request.URL.String() + "\n请检查服务器运行情况\n")
	}

	timeUse := time.Since(start).Nanoseconds() / 1e6
	defer res.Body.Close()

	body, err2 := ioutil.ReadAll(res.Body)
	if err2 != nil {
		return nil, err2
	}

	return &Response{
		Body:       body,
		Url:        r.request.URL.String(),
		StatusCode: res.StatusCode,
		Time:       int(timeUse),
	}, nil
}

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
