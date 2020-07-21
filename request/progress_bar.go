package request

import (
	"io"
	"io/ioutil"
	"net/http"

	api "github.com/guonaihong/gout/interface"
)

//上传进度条
type progressBar struct {
	callback   func(currBytes, totalBytes int)
	r          io.Reader
	currBytes  int
	totalBytes int
}

func (g *progressBar) Read(p []byte) (n int, err error) {
	n, err = g.r.Read(p)
	g.currBytes += n
	if n > 0 && g.callback != nil {
		g.callback(g.currBytes, g.totalBytes)
	}

	return
}

func (g *progressBar) ModifyRequest(req *http.Request) error {
	g.r = req.Body
	req.Body = ioutil.NopCloser(g)
	g.totalBytes = int(req.ContentLength)
	return nil
}

func ProgressBar(callback func(currBytes, totalBytes int)) api.RequestMiddler {
	return &progressBar{callback: callback}
}
