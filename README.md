## gout-middleware
![Go](https://github.com/antlabs/gout-middleware/workflows/Go/badge.svg)
[![codecov](https://codecov.io/gh/antlabs/gout-middleware/branch/master/graph/badge.svg)](https://codecov.io/gh/antlabs/gout)

gout请求和响应中间件项目
## 请求中间件
- [gzip](#gzip)
  - [请求body使用gzip压缩](#请求body使用gzip压缩)
  - [设置请求body大于一定字节数才压缩](#设置请求body大于一定字节数才压缩)
- [unzip](#unzip)
  - [解压缩body里面的gzip数据](#解压缩body里面的gzip数据)
- [upload 进度条](#upload进度条)
- [close 3xx自动跳转](#close-3xx自动跳转)
### gzip
#### 请求body使用gzip压缩
```go
import (
        "github.com/antlabs/gout-middleware/request"
        "github.com/guonaihong/gout"
)

func main() {
        gout.POST(":6666/compress").
                RequestUse(request.GzipCompress()).
                SetBody("hello world").
                Do()
}

```
#### 设置请求body大于一定字节数才压缩
```go
import (
	"github.com/antlabs/gout-middleware/request"
	"github.com/guonaihong/gout"
)

func main() {
	gout.POST(":6666/compress").
		RequestUse(request.GzipCompress(request.EnableGzipGreaterEqual(4))). //大于等于4个字节才压缩
		SetBody("hello world").
		Do()
}

```
### unzip
#### 解压缩body里面的gzip数据
```go
import (
	"github.com/antlabs/gout-middleware/request"
        "github.com/guonaihong/gout"
        "bytes"
)

func main() {
        var buf bytes.Buffer //假装buf里面有gzip数据
        gout.POST(":6666/compress").RequestUse(request.GzipDecompress()).SetBody(buf).Do()
}
```
### upload进度条
```go
package main

import (
        "bytes"
        "github.com/antlabs/gout-middleware/request"
        "github.com/guonaihong/gout"
)

func main() {
        gout.POST(":8080").RequestUse(request.ProgressBar(func(currBytes, totalBytes int) {

                fmt.Printf("%d:%d-->%f%%\n", currBytes, totalBytes, float64(currBytes)/float64(totalBytes))
        })).SetBody(strings.Repeat("1", 100000) /*构造大点的测试数据，这里换成真实业务数据*/).Do()
}

```
### close 3xx自动跳转
```go
package main

import (
	"github.com/antlabs/gout-middleware/request"
	"github.com/guonaihong/gout"
	"net/http"
)

func main() {
	c := &http.Client{}
	gout.New(c).GET(":8080/301").RequestUse(request.Close3xx(c, true)).Do()
}

```