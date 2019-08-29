package add

import (
	"encoding/json"
	"net/http"
	"strconv"
)

// Function to add two numbers without using an arithmetic operator
func IntAdd(a, b int64) int64 {
	for b != 0 {
		var carry = a & b
		a = a ^ b
		b = carry << 1
	}
	return a
}

func AddHandler(w http.ResponseWriter, r *http.Request){
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
	result := IntAdd(num1, num2)

	// resMap := make(map[string]int64)
	// resMap["result"] = result
	// w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	json.NewEncoder(w).Encode(result)
}
