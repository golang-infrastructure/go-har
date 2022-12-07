# HAR（HTTP Archive format）HTTP归档格式文件的Golang解析库



# 一、安装

```bash
go get -u github.com/golang-infrastructure/go-har
```

# 二、代码示例

```go
package main

import (
	"fmt"
	"github.com/golang-infrastructure/go-har"
	"os"
)

func main() {

	harFilePath := "./data/www.google.com.har"

	// 用法一：解析HAR文件的字节
	harFileBytes, err := os.ReadFile(harFilePath)
	if err != nil {
		fmt.Println(err)
		return
	}
	harFile, err := har.ParseHar(harFileBytes)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(harFile)

	// 用法二：给定HAR文件的路径解析它
	harFile002, err := har.ParseHarFile(harFilePath)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(harFile002)

}
```

