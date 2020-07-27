package request

import (
	"net/http"

	api "github.com/guonaihong/gout/interface"
)

type close3xx struct {
	close bool
	c     *http.Client
}

func (c *close3xx) ModifyRequest(req *http.Request) error {
	client := c.c
	if c.close == true {
		if client.CheckRedirect == nil {
			client.CheckRedirect = func(req *http.Request, via []*http.Request) error {
				return http.ErrUseLastResponse
			}
		}
		return nil
	}

	// c.close == false
	if client.CheckRedirect != nil {
		client.CheckRedirect = nil
	}
	return nil
}

// close 为true 则关闭3xx跳转功能
// close 为false 则不关闭3xx跳转功能
func Close3xx(c *http.Client, close bool) api.RequestMiddler {
	return &close3xx{c: c, close: close}
}
