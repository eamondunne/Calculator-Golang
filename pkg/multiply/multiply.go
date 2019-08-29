package multiply

import (
	"net/http"
	"strconv"
	"encoding/json"
	"github.com/eamondunne/calculator/pkg/add"
	"github.com/eamondunne/calculator/pkg/subtract"
)

// Function to multiply two numbers without using arithmetic operators
func IntMultiply(a, b int64) int64 {
	var product int64 = 0
	var neg_flag bool = false
	var additive_inverse int64

	// Check if either digit is 0
	if (a == 0 || b == 0){
		return 0
	}

	// If a is negative, convert to positive, invert flag and reconvert at end
	if (a < 0){
		neg_flag = !neg_flag
		additive_inverse = a
		// a = -a
		a = subtract.IntSubtract(a, additive_inverse)
		a = subtract.IntSubtract(a, additive_inverse)
	}

	// If b is negative, convert to positive, invert flag and reconvert at end
	if (b < 0){
		neg_flag = !neg_flag
		additive_inverse = b
		// b = -b
		b = subtract.IntSubtract(b, additive_inverse)
		b = subtract.IntSubtract(b, additive_inverse)
	}

	// a * b
	for a != 0 {
		if (a % 2 == 1) {
			product = add.IntAdd(product, b)
		}
		a >>= 1
		b <<= 1
	}

	// If flag is set, reconvert to negative value
	if(neg_flag){
		additive_inverse = product
		// product = -product
		product = subtract.IntSubtract(product, additive_inverse)
		product = subtract.IntSubtract(product, additive_inverse)
	}
	return product
}

func MultiplyHandler(w http.ResponseWriter, r *http.Request){
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
	result := IntMultiply(num1, num2)
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	json.NewEncoder(w).Encode(result)
}