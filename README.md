# go-version 
A package helps to get golang latest and all version go's  download url\file size\checksum\os\arch etc.
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
[中文](./README_CN.md)
