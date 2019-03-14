package handler

import (
	"fmt"
	"math"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

// Handler for prime request.
// It returns the json to client.
func Prime(c *gin.Context) {
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
		p := ResultPrime{}
		i := 2

		// Loop for get number of prime
		for {
			if IsPrime(i) {
				p.Primes = append(p.Primes, i)
			}

			if int(nInt) <= len(p.Primes) {
				break
			}

			i++
		}

		p.String = strings.Trim(strings.Join(strings.Fields(fmt.Sprint(p.Primes)), ", "), "[]")
		res.Data = p
	}

	c.JSON(http.StatusOK, res)
}

// Check the number is prime or not.
// Return the boolean, true for prime.
func IsPrime(val int) bool {
	for i := 2; i <= int(math.Floor(float64(val)/2)); i++ {
		if val%i == 0 {
			return false
		}
	}

	return val > 1
}
