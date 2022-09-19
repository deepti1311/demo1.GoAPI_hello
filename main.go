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

	//HTTP GET,POST,PUT,DELETE

	r.GET("/", getHello)
	r.POST("/", POSTHello)
	r.PUT("/", putHello)
	r.DELETE("/", deleteHello)

	//r.Run()  //run 8080
	r.Run(":8090")
	fmt.Println("server is running")

}

func hello(c *gin.Context) {

	c.String(http.StatusOK, "HEllo2 rest api")

}

func getHello(c *gin.Context) {

	c.String(http.StatusOK, "HEllo GET rest api")

}

func POSTHello(c *gin.Context) {

	c.String(http.StatusOK, "HEllo POST rest api")

}

func putHello(c *gin.Context) {

	c.String(http.StatusOK, "HEllo PUT rest api")

}

func deleteHello(c *gin.Context) {

	c.String(http.StatusOK, "HEllo Delete rest api")

}
