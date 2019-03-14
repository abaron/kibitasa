package main

import (
	"github.com/abaron/kibitasa/src/handler"
	"github.com/gin-gonic/gin"
)

func main() {

	// Set the router as the default one shipped with Gin
	r := gin.Default()

	// Set route for addition request
	r.GET("/sum/:a/:b", handler.Sum)

	// Set route for multiplication request
	r.GET("/multiple/:a/:b", handler.Multiple)

	// Set route for prime request
	r.GET("/prime/:n", handler.Prime)

	// Set route for fibonacci request
	r.GET("/fibonacci/:n", handler.Fibonacci)

	// listen and serve on 0.0.0.0:8080
	r.Run()
}
