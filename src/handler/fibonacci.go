package handler

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

var Fibo []int

// Handler for fibonacci request.
// It returns the json to client.
func Fibonacci(c *gin.Context) {
	// Parse parameter from url
	n := c.Param("n")

	// Set default response
	res := Response{
		Code:    200,
		Status:  "Ok",
		Message: "",
	}

	// Convert `n` to int64
	nInt, err := strconv.ParseInt(n, 0, 64)
	if err != nil || nInt < 1 {
		res.Code = 400
		res.Status = "Error"
		res.Message = InvalidParam
	}

	if res.Code == 200 {
		f := ResultFibonacci{}

		// Write recursive fibonacci number
		GetFibonacci(int(nInt))

		f.Fibonacci = Fibo
		f.String = strings.Trim(strings.Join(strings.Fields(fmt.Sprint(f.Fibonacci)), ", "), "[]")
		res.Data = f
	}

	c.JSON(http.StatusOK, res)
}

// Write recursive fibonacci number.
// Fibonacci started from 0 then 1.
// It will be writting the Fibo variable.
func GetFibonacci(val int) {
	if val <= len(Fibo) {
		return
	} else if len(Fibo) == 0 {
		Fibo = append(Fibo, 0)
	} else if len(Fibo) == 1 && val > 1 {
		Fibo = append(Fibo, 1)
	} else {
		Fibo = append(Fibo, Fibo[len(Fibo)-1]+Fibo[len(Fibo)-2])
	}

	GetFibonacci(val)
}
