package request

import (
	"bytes"
	stdgzip "compress/gzip"
	"fmt"
	"testing"

	"github.com/guonaihong/gout"
	"github.com/stretchr/testify/assert"
)

func Test_GzipDecompress(t *testing.T) {

	buf := &bytes.Buffer{}

	w := stdgzip.NewWriter(buf)
	w.Write([]byte(testGzipData))
	w.Close()

	fmt.Printf("##buf.len:%d:%x\n", buf.Len(), buf.Bytes())
	ts := createNotDeCompressServer()
	var got string
	err := gout.POST(ts.URL).RequestUse(GzipDecompress()).SetBody(buf).BindBody(&got).Do()
	assert.NoError(t, err)
	assert.Equal(t, got, testGzipData)
}
