package multiply_test

import (
	. "github.com/eamondunne/calculator/pkg/multiply"
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestIntMultiply(t *testing.T) {
	var tests = []struct {
		input1    int64
		input2		int64
		expected int64
	}{
		{6, 12, 72},
		{0, 0, 0},
		{0, 512346, 0},
		{1000, -0, 0},
		{-100, -2000, 200000},
	}
	
	for _, test := range tests {
		assert.Equal(t, test.expected, IntMultiply(test.input1, test.input2))
	}
}
