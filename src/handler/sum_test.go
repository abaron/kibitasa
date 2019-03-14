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

// Simple addition test
func TestSum(t *testing.T) {
	router := gin.New()
	router.GET("/test/sum/:a/:b", handler.Sum)

	req, _ := http.NewRequest("GET", "/test/sum/3/5", nil)
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	assert.Equal(
		t,
		"{\"code\":200,\"status\":\"Ok\",\"message\":\"\",\"data\":{\"result\":8}}",
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
		if k.String() == "result" {
			assert.Equal(t, float64(8), s.MapIndex(k).Interface())
		}
	}
}

// Test to get negative number
func TestSumResultMin(t *testing.T) {
	router := gin.New()
	router.GET("/test/sum/:a/:b", handler.Sum)

	req, _ := http.NewRequest("GET", "/test/sum/-3/-5", nil)
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	assert.Equal(
		t,
		"{\"code\":200,\"status\":\"Ok\",\"message\":\"\",\"data\":{\"result\":-8}}",
		resp.Body.String(),
	)
}

// Addition decimal / float number
func TestSumFloat(t *testing.T) {
	router := gin.New()
	router.GET("/test/sum/:a/:b", handler.Sum)

	req, _ := http.NewRequest("GET", "/test/sum/3.0/5.5", nil)
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	assert.Equal(
		t,
		"{\"code\":200,\"status\":\"Ok\",\"message\":\"\",\"data\":{\"result\":8.5}}",
		resp.Body.String(),
	)
}

// Test invalid first parameter
func TestSumInvalidFirstParam(t *testing.T) {
	router := gin.New()
	router.GET("/test/sum/:a/:b", handler.Sum)

	req, _ := http.NewRequest("GET", "/test/sum/a/5", nil)
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	assert.Equal(
		t,
		"{\"code\":400,\"status\":\"Error\",\"message\":\"Invalid Parameter(s)\",\"data\":null}",
		resp.Body.String(),
	)
}

// Test invalid second parameter
func TestSumInvalidSecondParam(t *testing.T) {
	router := gin.New()
	router.GET("/test/sum/:a/:b", handler.Sum)

	req, _ := http.NewRequest("GET", "/test/sum/5/5a", nil)
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	assert.Equal(
		t,
		"{\"code\":400,\"status\":\"Error\",\"message\":\"Invalid Parameter(s)\",\"data\":null}",
		resp.Body.String(),
	)
}
