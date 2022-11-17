package main

import (
	"gee"
	"net/http"
)

func main() {
	r := gee.New()

	//r.GET("/", func(w http.ResponseWriter, r *http.Request) {
	//	fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)
	//})
	//r.GET("/hello", func(w http.ResponseWriter, r *http.Request) {
	//	for k, v := range r.Header {
	//		fmt.Fprintf(w, "Header[%q] = %q\n", k, v)
	//	}
	//})
	r.GET("/", func(c *gee.Context) {
		c.HTML(http.StatusOK, "<h1>Hello, Gee.</h1>")
	})
	r.GET("/hello", func(c *gee.Context) {
		c.String(http.StatusOK, "Hello %s, you are at %s\n", c.Query("name"), c.Path)
	})
	r.POST("/login", func(c *gee.Context) {
		c.JSON(http.StatusOK, gee.H{
			"username": c.PostForm("username"),
			"password": c.PostForm("password"),
		})
	})
	r.RUN(":9999")
}
