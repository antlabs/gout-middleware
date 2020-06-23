## gout-middleware
gout请求和响应中间件项目

### 请求body使用gzip压缩
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