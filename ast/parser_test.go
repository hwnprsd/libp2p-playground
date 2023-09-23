package ast

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	valueMap1 = map[int]bool{
		1: true,
		2: true,
		3: true,
	}

	valueMap2 = map[int]bool{
		1: true,
		2: false,
		3: false,
	}

	valueMap3 = map[int]bool{
		1: false,
		2: false,
		3: false,
	}
)

func Test_Parser(t *testing.T) {
	v1 := ValidateExpression("[1] OR [2]", valueMap1)
	assert.True(t, v1)

	v2 := ValidateExpression("[1] OR [2]", valueMap2)
	assert.True(t, v2)

	v3 := ValidateExpression("[1] OR [2]", valueMap3)
	assert.False(t, v3)
}
