package divide_test

import (
	. "github.com/eamondunne/calculator/pkg/divide"
	"testing"
	"github.com/stretchr/testify/assert"
	"fmt"
)
func TestIntDivide(t *testing.T) {
	var tests = []struct {
		input1    int64
		input2		int64
		expected  int64
	}{
		{100, 2, 50},
		{10, 0, 0},
		{6, 3, 2},
		{14, 2, 7},
		{3, 4, 0},
		{7, 8, 0},
		{-4, -1, 4},
		{-16, 2, -8},
	}
	
	for _, test := range tests {
		res, err := IntDivide(test.input1, test.input2)
		if (err != nil){
			fmt.Println(err)
		}
		assert.Equal(t,test.expected, res)
	}
}