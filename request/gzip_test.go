package request

import (
	"bytes"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	stdgzip "compress/gzip"

	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"
	"github.com/guonaihong/gout"
	"gopkg.in/go-playground/assert.v1"
)

const (
	testGzipData = "123456789abcdefgh"
)

func createNotDeCompressServer() *httptest.Server {
	r := gin.New()

	//r.Use(gzip.Gzip(gzip.DefaultCompression, gzip.WithDecompressFn(gzip.DefaultDecompressHandle)))
	r.POST("/", func(c *gin.Context) {
		io.Copy(c.Writer, c.Request.Body)
	})

	return httptest.NewServer(http.HandlerFunc(r.ServeHTTP))
}

func createDeCompressServer() *httptest.Server {
	r := gin.New()

	r.Use(gzip.Gzip(gzip.DefaultCompression, gzip.WithDecompressFn(gzip.DefaultDecompressHandle)))
	r.POST("/", func(c *gin.Context) {
		io.Copy(c.Writer, c.Request.Body)
	})

	return httptest.NewServer(http.HandlerFunc(r.ServeHTTP))
}

func TestGzipCompress(t *testing.T) {
	// 客户端压缩 + 服务不解压缩
	got := []byte{}
	ts := createNotDeCompressServer()
	gout.POST(ts.URL).RequestUse(GzipCompress()).SetBody(testGzipData).BindBody(&got).Do()

	buf := &bytes.Buffer{}

	w := stdgzip.NewWriter(buf)
	w.Write([]byte(testGzipData))
	w.Close()

	assert.Equal(t, got, buf.Bytes())

	// 客户端压缩 + 服务解压缩
	got = []byte{}
	ts = createDeCompressServer()
	gout.POST(ts.URL).RequestUse(GzipCompress()).SetBody(testGzipData).BindBody(&got).Do()

	buf = &bytes.Buffer{}

	w = stdgzip.NewWriter(buf)
	w.Write([]byte(testGzipData))
	w.Close()

	assert.Equal(t, got, []byte(testGzipData))
}
