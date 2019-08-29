package add_test

import (
	. "github.com/eamondunne/calculator/add"
	"testing"
	"github.com/stretchr/testify/assert"
)
func TestIntAdd(t *testing.T) {
	var tests = []struct {
		input1    int64
		input2		int64
		expected  int64
	}{
		{2, 3, 5},
		{0, 0, 0},
		{-100, 200, 100},
	}
	
	for _, test := range tests {
		assert.Equal(t, test.expected, IntAdd(test.input1, test.input2))
	}
}