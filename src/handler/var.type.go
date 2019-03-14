package handler

var (
	InvalidParam = "Invalid Parameter(s)" // Error message
)

// Default response for every valid request
// As paramater for gin.Context.JSON
type Response struct {
	Code    int         `json:"code"`    // 200 - 500
	Status  string      `json:"status"`  // Ok or Error
	Message string      `json:"message"` // Error message
	Data    interface{} `json:"data"`    // Result data
}

// Result data for addition and multiplication request
type Result struct {
	Result interface{} `json:"result"` // Result data
}

// Result data for prime request
type ResultPrime struct {
	Primes []int  `json:"primes"` // Result data (array)
	String string `json:"string"` // Result data joined to string
}

// Result data for fibonacci request
type ResultFibonacci struct {
	Fibonacci []int  `json:"fibonacci"` // Result data (array)
	String    string `json:"string"`    // Result data joined to string
}
