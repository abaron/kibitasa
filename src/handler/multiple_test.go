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

// Simple multiple test
func TestMultiple(t *testing.T) {
	router := gin.New()
	router.GET("/test/multiple/:a/:b", handler.Multiple)

	req, _ := http.NewRequest("GET", "/test/multiple/3/5", nil)
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	assert.Equal(
		t,
		"{\"code\":200,\"status\":\"Ok\",\"message\":\"\",\"data\":{\"result\":15}}",
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
			assert.Equal(t, float64(15), s.MapIndex(k).Interface())
		}
	}
}

// Test to get negative number
func TestMultipleNegative(t *testing.T) {
	router := gin.New()
	router.GET("/test/multiple/:a/:b", handler.Multiple)

	req, _ := http.NewRequest("GET", "/test/multiple/3/-5", nil)
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	assert.Equal(
		t,
		"{\"code\":200,\"status\":\"Ok\",\"message\":\"\",\"data\":{\"result\":-15}}",
		resp.Body.String(),
	)
}

// Addition decimal / float number
func TestMultipleFloat(t *testing.T) {
	router := gin.New()
	router.GET("/test/multiple/:a/:b", handler.Multiple)

	req, _ := http.NewRequest("GET", "/test/multiple/3.0/5.5", nil)
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	assert.Equal(
		t,
		"{\"code\":200,\"status\":\"Ok\",\"message\":\"\",\"data\":{\"result\":16.5}}",
		resp.Body.String(),
	)
}

// Test invalid first parameter
func TestMultipleInvalidFirstParam(t *testing.T) {
	router := gin.New()
	router.GET("/test/multiple/:a/:b", handler.Multiple)

	req, _ := http.NewRequest("GET", "/test/multiple/a/5", nil)
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	assert.Equal(
		t,
		"{\"code\":400,\"status\":\"Error\",\"message\":\"Invalid Parameter(s)\",\"data\":null}",
		resp.Body.String(),
	)
}

// Test invalid second parameter
func TestMultipleInvalidSecondParam(t *testing.T) {
	router := gin.New()
	router.GET("/test/multiple/:a/:b", handler.Multiple)

	req, _ := http.NewRequest("GET", "/test/multiple/5/5a", nil)
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	assert.Equal(
		t,
		"{\"code\":400,\"status\":\"Error\",\"message\":\"Invalid Parameter(s)\",\"data\":null}",
		resp.Body.String(),
	)
}
