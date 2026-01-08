Basic Go WebServer using the Gin framework

```go
package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		//c.String(200, "¡Hola, Mundo!")
		c.JSON(200, gin.H{
			"message": "¡Hola, mundo!",
		})
	})

	r.Run(":8080")
}
```