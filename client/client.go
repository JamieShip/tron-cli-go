package client

import (
	"net"
	"net/http"
	"time"

	"github.com/go-resty/resty/v2"
)

type TronClient struct {
	c *resty.Client
}

func NewTronClient(rpc_endpoint string) *TronClient {
	_c := resty.New().SetBaseURL(rpc_endpoint).SetHeader("Content-Type", "application/json").
		SetTransport(&http.Transport{
			MaxIdleConns:        1000,
			MaxIdleConnsPerHost: 1000,
			IdleConnTimeout:     80 * time.Second,
			DialContext: (&net.Dialer{
				Timeout:   30 * time.Second,
				KeepAlive: 30 * time.Second,
			}).DialContext,
		}).
		SetTimeout(30 * time.Second)
	return &TronClient{
		c: _c,
	}
}

func (t *TronClient) WithTransport(transport *http.Transport) *TronClient {
	t.c.SetTransport(transport)
	return t
}
