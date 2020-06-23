package request

import (
	"bytes"
	"compress/gzip"
	api "github.com/guonaihong/gout/interface"
	"io"
	"io/ioutil"
	"net/http"
)

type gzipCompress struct{}

func (g *gzipCompress) ModifyRequest(req *http.Request) error {
	// 如果已经有一种编码格式，不会生效
	if len(req.Header.Get("Content-Encoding")) > 0 {
		return nil
	}

	buf := &bytes.Buffer{}

	w := gzip.NewWriter(buf)
	body, err := req.GetBody()
	if err != nil {
		return nil
	}

	io.Copy(w, body)
	w.Close()

	if req.ContentLength > 0 {
		req.ContentLength = int64(buf.Len())
	}

	req.Body = ioutil.NopCloser(buf)
	req.Header.Set("Content-Encoding", "gzip")
	return nil
}
func GzipCompress() api.RequestMiddler {
	return &gzipCompress{}
}
