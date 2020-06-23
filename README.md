## gout-middleware
![Go](https://github.com/antlabs/gout-middleware/workflows/Go/badge.svg)
[![codecov](https://codecov.io/gh/antlabs/gout-middleware/branch/master/graph/badge.svg)](https://codecov.io/gh/antlabs/gout)

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
