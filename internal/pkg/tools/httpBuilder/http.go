package httpBuilder

import (
	"bytes"
	"fmt"
	"golang.org/x/net/proxy"
	"io/ioutil"
	"net/http"
	"net/url"
)

type Req struct {
	Method  string
	Headers map[string]string
	Url     string
	Query   map[string]string
	Body    []byte
}

func NewRequest() *Req {
	return &Req{}
}

func (req *Req) Do(proxy string) ([]byte, error) {
	val := url.Values{}
	for key, value := range req.Query {
		val.Add(key, value)
	}
	httpReq, err := http.NewRequest(req.Method, fmt.Sprintf("%s?%s", req.Url, val.Encode()),
		bytes.NewReader(req.Body))
	if err != nil {
		return nil, err
	}
	client := &http.Client{}
	if proxy != "" {
		client, err = req.setSocksProxy(proxy)
		if err != nil {
			return nil, err
		}
	}
	resp, err := client.Do(httpReq)
	if err != nil {
		return nil, err
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return body, nil
}

func (req *Req) setSocksProxy(addr string) (*http.Client, error) {
	// create a socks5 dialer
	u, err := url.Parse(addr)
	if err != nil {
		return nil, err
	}
	pwd, _ := u.User.Password()
	dialer, err := proxy.SOCKS5("tcp", fmt.Sprintf("%s:%s", u.Hostname(), u.Port()), &proxy.Auth{
		User:     u.User.Username(),
		Password: pwd,
	}, proxy.Direct)
	if err != nil {
		return nil, err
	}
	// setup a http client
	httpTransport := &http.Transport{}
	httpClient := &http.Client{Transport: httpTransport}
	// set our socks5 as the dialer
	httpTransport.Dial = dialer.Dial
	return httpClient, nil
}
