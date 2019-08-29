package subtract_test

import (
	. "github.com/eamondunne/calculator/subtract"
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestIntSubtract(t *testing.T) {
	var tests = []struct {
		input1    int64
		input2		int64
		expected  int64
	}{
		{5, 3, 2},
		{0, 0, 0},
		{-100, 200, -300},
	}
	
	for _, test := range tests {
		assert.Equal(t, test.expected, IntSubtract(test.input1, test.input2))
	}
}
