package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {

	r := gin.Default()

	r.GET("/hello", func(c *gin.Context) {
		c.String(http.StatusOK, "hello rets api")
	})

	r.GET("hello2", hello)

	//r.Run()  //run 8080
	r.Run()
	fmt.Println("server is running")

}

func hello(c *gin.Context) {

	c.String(http.StatusOK, "HEllo2 rest api")

}
