package request

import (
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/guonaihong/gout"
	"github.com/stretchr/testify/assert"
)

type readWait struct {
	r io.Reader
}

func (r *readWait) Read(p []byte) (n int, err error) {
	n, err = r.r.Read(p[:512])
	return
}

func createProgressBarServer() *httptest.Server {
	r := gin.New()
	r.POST("/", func(c *gin.Context) {
		r := &readWait{r: c.Request.Body}
		_, _ = ioutil.ReadAll(r)

		c.String(200, "hello")
	})

	return httptest.NewServer(http.HandlerFunc(r.ServeHTTP))
}

func TestProgressBar(t *testing.T) {
	ts := createProgressBarServer()
	call := false
	gout.POST(ts.URL).RequestUse(ProgressBar(func(currBytes, totalBytes int) {
		call = true
		//fmt.Printf("%d:%d\n", currBytes, totalBytes)
	})).SetBody(strings.Repeat("1", 100000)).Do()

	assert.True(t, call)
}
