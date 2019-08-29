package divide

import (
	"fmt"
	"errors"
	"net/http"
	"strconv"
	"encoding/json"
	"github.com/eamondunne/calculator/pkg/subtract"
	"github.com/eamondunne/calculator/pkg/add"
)

func IntDivide(a, b int64) (result int64, err error) {
	var sign int64 = 1
	// var temp, quotient int64 = 0, 0
	var quotient int64 = 0
	var neg_flag bool = false
	var additive_inverse int64
	// cannot divide by 0
	if (b == 0){
		return 0, errors.New("cannot divide by zero")
	}

	// if a is negative, abs(a) + set flag
	if (a < 0){
		neg_flag = !neg_flag
		additive_inverse = a
		// a = -a
		a = subtract.IntSubtract(a, additive_inverse)
		a = subtract.IntSubtract(a, additive_inverse)
	}
	// if b is negative, abs(b) + set flag
	if (b < 0){
		neg_flag = !neg_flag
		additive_inverse = b
		// b = -b
		b = subtract.IntSubtract(b, additive_inverse)
		b = subtract.IntSubtract(b, additive_inverse)
	}

	// Calculate Quotient of a / b
	for {
		a = subtract.IntSubtract(a, b)
		if (a >= 0){
			quotient = add.IntAdd(quotient, 1)
		} else {
			break
		}
	}
	if(neg_flag){
		additive_inverse = quotient
		// quotient = -quotient
		quotient = subtract.IntSubtract(quotient, additive_inverse)
		quotient = subtract.IntSubtract(quotient, additive_inverse)
	}
	return (sign * quotient), nil;
}

func DivideHandler(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
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
	result, err := IntDivide(num1, num2)
	if(err != nil){
		fmt.Println("cannot divide by 0")
		json.NewEncoder(w).Encode("Infinity")
	} else {
		json.NewEncoder(w).Encode(result)
	}
}