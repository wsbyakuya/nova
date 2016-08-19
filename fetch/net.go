package fetch

import (
	"errors"
	"io"
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

func (r *NovaRequest) AddUserAgent(userAgent string) {
	r.request.Header.Add("User-Agent", userAgent)
}

func (r *NovaRequest) SetBody(body string) {
	r.request.Body = NewBodyReadCloser(body)
}

func (r *NovaRequest) SetCookies(cookies map[string]string) {
	if len(cookies) == 0 {
		return
	}
	for k, v := range cookies {
		r.AddCookie(k, v)
	}
}

func (r *NovaRequest) SetHeader(header map[string]string) {
	if len(header) == 0 {
		return
	}
	for k, v := range header {
		r.request.Header.Set(k, v)
	}
}

func (r *NovaRequest) Do() (response *Response, err error) {
	start := time.Now()

	res, err := http.DefaultClient.Do(r.request)
	if err != nil {
		return nil, errors.New("无法获取连接：" + r.request.URL.String() + "\n请检查服务器运行情况\n")
	}

	timeUse := time.Since(start).Nanoseconds() / 1e6 //单位：ms
	defer res.Body.Close()

	body, err2 := ioutil.ReadAll(res.Body)
	if err2 != nil {
		return nil, err2
	}

	return &Response{
		Body:       body,
		Url:        r.request.URL.String(),
		Status:     res.Status,
		StatusCode: res.StatusCode,
		Time:       int(timeUse),
	}, nil
}

type Response struct {
	Body       []byte
	Url        string
	Status     string
	StatusCode int
	Time       int `ms`
}

func HttpGet(urlStr string) (response *Response, err error) {
	req, err := NewNovaRequest("GET", urlStr)
	if err != nil {
		return nil, err
	}
	res, err2 := req.Do()
	return res, err2
}

func HttpPost(url, body string) (response *Response, err error) {
	req, err := NewNovaRequest("POST", url)
	if err != nil {
		return nil, err
	}
	req.SetBody(body)
	res, err2 := req.Do()
	return res, err2
}

//BodyReadCloser实现io.ReadCloser的结构体，用于POST方法中读出字符串
type BodyReadCloser struct {
	s string
	i int
}

func NewBodyReadCloser(body string) *BodyReadCloser {
	return &BodyReadCloser{
		s: body,
		i: 0,
	}
}

func (p *BodyReadCloser) Read(b []byte) (n int, err error) {
	if len(p.s) == 0 {
		return 0, nil
	}
	if p.i >= len(p.s) {
		return 0, io.EOF
	}
	n = copy(b, []byte(p.s)[p.i:])
	p.i += n
	return n, nil
}

func (p *BodyReadCloser) Close() error {
	return nil
}
