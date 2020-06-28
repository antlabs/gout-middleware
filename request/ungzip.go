package request

import (
	"bytes"
	"compress/gzip"
	"io"
	"io/ioutil"
	"net/http"

	api "github.com/guonaihong/gout/interface"
)

type gzipDecompress struct{}

var gzipHead = []byte{0x1f, 0x8b}

func (g *gzipDecompress) ModifyRequest(req *http.Request) error {
	var saveGzip bytes.Buffer
	io.Copy(&saveGzip, req.Body)

	if saveGzip.Len() == 0 || saveGzip.Len() < 2 {
		return nil
	}

	// https://wiki.fileformat.com/compression/gz/
	// gzip header 0x1f 0x8b 是gzip的魔数id
	if !bytes.Equal(saveGzip.Bytes()[:2], gzipHead) {
		return nil
	}

	r, err := gzip.NewReader(&saveGzip)
	if err != nil {
		return nil
	}

	var raw bytes.Buffer
	io.Copy(&raw, r)
	r.Close()

	req.Body = ioutil.NopCloser(&raw)
	req.ContentLength = int64(raw.Len())
	return nil
}

func GzipDecompress() api.RequestMiddler {
	return &gzipDecompress{}
}
