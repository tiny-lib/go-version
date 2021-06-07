# go-version
一个用来获取Golang官方版本信息的包。可获取包括版本号、下载链接、文件校验等信息。
下面是一个简单的例子
a little example
```go
package main

import (
	"fmt"
	"log"

	goversion "github.com/czyt/go-version"
)

func main() {
	downLoader := goversion.NewDownLoader("https://golang.google.cn")
	versions, err := downLoader.GetFeaturedDownload()
	if err != nil {
		log.Fatal(err)
	}
	for _, v := range versions {
		fmt.Println(v.DownloadUrl)
	}
}

```
[English](./README.md)
