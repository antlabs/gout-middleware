package request

import (
	"bytes"
	stdgzip "compress/gzip"
	"testing"

	"github.com/guonaihong/gout"
	"github.com/stretchr/testify/assert"
)

// 测试数据小于2
func Test_GzipDecompress_SmallData(t *testing.T) {

	buf := &bytes.Buffer{}

	buf.Write([]byte("1"))

	ts := createNotDeCompressServer()
	var got string
	err := gout.POST(ts.URL).RequestUse(GzipDecompress()).SetBody(buf).BindBody(&got).Do()
	assert.NoError(t, err)
	assert.Equal(t, got, "1")
}

// 测试数据不是gzip
func Test_GzipDecompress_NotGzip(t *testing.T) {

	buf := &bytes.Buffer{}

	buf.Write([]byte(testGzipData))

	ts := createNotDeCompressServer()
	var got string
	err := gout.POST(ts.URL).RequestUse(GzipDecompress()).SetBody(buf).BindBody(&got).Do()
	assert.NoError(t, err)
	assert.Equal(t, got, testGzipData)
}

// 测试数据是gzip
func Test_GzipDecompress(t *testing.T) {

	buf := &bytes.Buffer{}

	w := stdgzip.NewWriter(buf)
	w.Write([]byte(testGzipData))
	w.Close()

	ts := createNotDeCompressServer()
	var got string
	err := gout.POST(ts.URL).RequestUse(GzipDecompress()).SetBody(buf).BindBody(&got).Do()
	assert.NoError(t, err)
	assert.Equal(t, got, testGzipData)
}
