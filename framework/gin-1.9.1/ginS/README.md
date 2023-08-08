# Gin Default Server

This is API experiment for Gin.

```go
package main

import (
	"ASPGo/framework"
	"ASPGo/framework/gin-1.9.1/ginS"
)

func main() {
	ginS.GET("/", func(c *gin.Context) { c.String(200, "Hello World") })
	ginS.Run()
}
```
