package main

import (
	"net/http"
	"vegin"
)

func main() {
	r := vegin.New()
	r.GET("/", func(c *vegin.Context) {
		c.HTML(http.StatusOK, "<h1>Hello vegin</h1>")
	})

	r.GET("/hello", func(c *vegin.Context) {
		// expect /hello?name=veginktutu
		c.String(http.StatusOK, "hello %s, you're at %s\n", c.Query("name"), c.Path)
	})

	r.GET("/hello/:name", func(c *vegin.Context) {
		// expect /hello/veginktutu
		c.String(http.StatusOK, "hello %s, you're at %s\n", c.Param("name"), c.Path)
	})

	r.GET("/assets/*filepath", func(c *vegin.Context) {
		c.JSON(http.StatusOK, vegin.H{"filepath": c.Param("filepath")})
	})

	r.Run(":9999")
}
