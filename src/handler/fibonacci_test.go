package handler_test

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	"github.com/abaron/kibitasa/src/handler"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

// Simple fibonacci test
func TestFibonacci(t *testing.T) {
	router := gin.New()
	router.GET("/test/fibonacci/:n", handler.Fibonacci)

	req, _ := http.NewRequest("GET", "/test/fibonacci/3", nil)
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	assert.Equal(
		t,
		"{\"code\":200,\"status\":\"Ok\",\"message\":\"\",\"data\":{\"fibonacci\":[0,1,1],\"string\":\"0, 1, 1\"}}",
		resp.Body.String(),
	)

	var response handler.Response
	bytes := []byte(resp.Body.String())

	err := json.Unmarshal(bytes, &response)
	if err != nil {
		fmt.Println("Error Unmarshal TestSum")
	}

	s := reflect.ValueOf(response.Data)
	for _, k := range s.MapKeys() {
		if k.String() == "string" {
			assert.Equal(t, "0, 1, 1", s.MapIndex(k).Interface().(string))
		}
	}
}

// Test invalid parameter
func TestFibonacciInvalidParam(t *testing.T) {
	router := gin.New()
	router.GET("/test/fibonacci/:n", handler.Fibonacci)

	req, _ := http.NewRequest("GET", "/test/fibonacci/-3", nil)
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	assert.Equal(
		t,
		"{\"code\":400,\"status\":\"Error\",\"message\":\"Invalid Parameter(s)\",\"data\":null}",
		resp.Body.String(),
	)
}
