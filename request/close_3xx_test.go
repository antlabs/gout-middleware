package request

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/guonaihong/gout"
	"github.com/stretchr/testify/assert"
)

func createClose302() *httptest.Server {
	r := gin.New()
	r.GET("/302", func(c *gin.Context) {
		c.String(200, "302")
	})

	return httptest.NewServer(http.HandlerFunc(r.ServeHTTP))
}

func createClose301(url string) *httptest.Server {
	r := gin.New()
	r.GET("/301", func(c *gin.Context) {
		c.Redirect(302, url+"/302")
	})

	return httptest.NewServer(http.HandlerFunc(r.ServeHTTP))
}

func Test_Close3xx_True(t *testing.T) {
	ts := createClose301("")
	c := &http.Client{}
	got := ""
	err := gout.New(c).GET(ts.URL + "/301").RequestUse(Close3xx(c, true)).BindBody(&got).Do()
	assert.NoError(t, err)
	assert.Equal(t, got, `<a href="/302">Found</a>.

`)
}

func Test_Close3xx_False(t *testing.T) {
	ts302 := createClose302()
	ts := createClose301(ts302.URL)
	c := &http.Client{}
	got := ""
	err := gout.New(c).GET(ts.URL + "/301").RequestUse(Close3xx(c, false)).BindBody(&got).Do()
	assert.NoError(t, err)
	assert.Equal(t, got, "302")
}

func Test_Close3xx_False2(t *testing.T) {
	ts302 := createClose302()
	ts := createClose301(ts302.URL)
	c := &http.Client{}
	got := ""
	err := gout.New(c).GET(ts.URL+"/301").RequestUse(Close3xx(c, true), Close3xx(c, false)).BindBody(&got).Do()
	assert.NoError(t, err)
	assert.Equal(t, got, "302")
}
