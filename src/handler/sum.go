package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// Handler for addition request.
// It returns the json to client.
func Sum(c *gin.Context) {
	// Parse parameters from url
	a := c.Param("a")
	b := c.Param("b")

	// Set default response
	res := Response{
		Code:    200,
		Status:  "Ok",
		Message: "",
	}

	// Convert `a` to int64
	aInt, err := strconv.ParseFloat(a, 0)
	if err != nil {
		res.Code = 400
		res.Status = "Error"
		res.Message = InvalidParam
	}

	// Convert `b` to int64
	bInt, err := strconv.ParseFloat(b, 0)
	if err != nil {
		res.Code = 400
		res.Status = "Error"
		res.Message = InvalidParam
	}

	if res.Code == 200 {
		res.Data = Result{Result: aInt + bInt}
	}

	c.JSON(http.StatusOK, res)
}
