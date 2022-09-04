package main

import (
	"fmt"
	"net/http"

	"lcy-faas/function_module"
	r "lcy-faas/router"

	"github.com/gin-gonic/gin"
)

var router r.Router

func main() {

	router = r.Router{
		Paths: map[string]function_module.FunctionModule{},
	}

	r := gin.Default()

	r.GET("/function", func(c *gin.Context) {
		html := "<h1>Functions</h1>"
		for key, value := range router.Paths {
			html += fmt.Sprint("<a href=\"", key, "\">", value.Name, "</a><br>")
		}
		c.Header("Content-Type", "text/html; charset=utf-8")
		c.String(200, html)
	})

	r.POST("/function", func(c *gin.Context) {
		var body function_module.FunctionModule
		if err := c.BindJSON(&body); err != nil {
			c.JSON(http.StatusServiceUnavailable, gin.H{
				"message": "error",
			})
		}
		err := body.Build()
		if err != nil {
			c.JSON(http.StatusServiceUnavailable, gin.H{
				"message": "error:" + err.Error(),
			})
		}
		router.Insert(body.Path, body)
	})

	r.DELETE("/function", func(c *gin.Context) {
		var body function_module.FunctionModule
		if err := c.BindJSON(&body); err != nil {
			c.JSON(http.StatusServiceUnavailable, gin.H{
				"message": "error",
			})
		}
		//
		router.Delete(body.Path)
	})

	r.GET("/function/:path", func(c *gin.Context) {
		path := c.Param("path")
		funcModule, ok := router.Find(path)
		if !ok {
			c.JSON(http.StatusNotImplemented, gin.H{
				"message": "not found",
			})
			return
		}
		resp, err := funcModule.Run()
		if err != nil {
			c.JSON(http.StatusFailedDependency, gin.H{
				"function_error": err.Error(),
			})
		}
		html := fmt.Sprintf("<h2>function_name:%v</h2><h3>response_data:%v</h3>", funcModule.Name, resp)
		c.Header("Content-Type", "text/html; charset=utf-8")
		c.String(http.StatusOK, html)
	})

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
