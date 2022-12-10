package request

import (
	"bytes"
	"compress/gzip"
	"io"
	"io/ioutil"
	"net/http"

	"github.com/guonaihong/gout/middler"
)

// 大于或等于EnableGzipMore 字节数的才压缩
type EnableGzipGreaterEqual int

type gzipCompress struct {
	enableGzipGreaterEqual int
}

func (g *gzipCompress) ModifyRequest(req *http.Request) error {
	// 如果已经有一种编码格式，不会生效
	if len(req.Header.Get("Content-Encoding")) > 0 {
		return nil
	}

	if req.ContentLength == 0 {
		return nil
	}

	// TODO 如果http client可以使用chunk 方式发送数据
	// 就不能使用req.ContentLength作为数据长度
	if g.enableGzipGreaterEqual > 0 {
		if req.ContentLength < int64(g.enableGzipGreaterEqual) {
			return nil
		}
	}

	buf := &bytes.Buffer{}

	w := gzip.NewWriter(buf)

	io.Copy(w, req.Body)
	w.Close()

	if req.ContentLength > 0 {
		req.ContentLength = int64(buf.Len())
	}

	req.Body = ioutil.NopCloser(buf)
	req.Header.Set("Content-Encoding", "gzip")
	return nil
}

func GzipCompress(args ...interface{}) middler.RequestMiddler {

	compress := &gzipCompress{}

	for _, a := range args {
		switch a.(type) {
		case EnableGzipGreaterEqual:
			compress.enableGzipGreaterEqual = int(a.(EnableGzipGreaterEqual))
		}
	}

	return compress
}
