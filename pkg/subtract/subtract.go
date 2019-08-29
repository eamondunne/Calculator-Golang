package subtract

import (
	"encoding/json"
	"net/http"
	"strconv"
)

// Function to subtract two numbers using only logical operators
func IntSubtract(a, b int64) int64 {
	for b != 0 {
		var borrow = (^a) & b
		a = a ^ b
		b = borrow << 1
	}
	return a
}

func SubtractHandler(w http.ResponseWriter, r *http.Request){
	// Get URL Params
	urlParams := r.URL.Query()
	json_num1, ok1 := urlParams["num1"]
	json_num2, ok2 := urlParams["num2"]
	if(!ok1 || !ok2){
		//TODO: Handle Error
	}
	num1, err1 := strconv.ParseInt(json_num1[0], 10, 64) 
	num2, err2 := strconv.ParseInt(json_num2[0], 10, 64) 
	if (err1 != nil || err2 != nil){
		// TODO: Handle Error
	}
	result := IntSubtract(num1, num2)
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	json.NewEncoder(w).Encode(result)
}