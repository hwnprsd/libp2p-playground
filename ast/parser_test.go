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

	valueMap4 = map[int]bool{
		1: false,
		2: false,
		3: true,
	}

	// New valueMaps
	valueMap5 = map[int]bool{
		1: true,
		2: false,
		3: true,
		4: true,
		5: false,
	}
	valueMap6 = map[int]bool{
		1: false,
		2: true,
		3: false,
		4: false,
		5: true,
	}
)

func Test_simpleOr(t *testing.T) {
	v1 := ValidateExpression("[1] OR [2]", valueMap1)
	v2 := ValidateExpression("[1] OR [2]", valueMap2)
	v3 := ValidateExpression("[1] OR [2]", valueMap3)
	assert.True(t, v1)
	assert.True(t, v2)
	assert.False(t, v3)
}

func Test_simpleAnd(t *testing.T) {
	v1 := ValidateExpression("[1] AND [2]", valueMap1)
	v2 := ValidateExpression("[1] AND [2]", valueMap2)
	v3 := ValidateExpression("[1] AND [2]", valueMap3)
	assert.True(t, v1)
	assert.False(t, v2)
	assert.False(t, v3)
}

func Test_nestedLogic(t *testing.T) {
	v1 := ValidateExpression("([1] AND [2]) OR [3]", valueMap4)
	v2 := ValidateExpression("([1] OR [2]) AND [3]", valueMap4)
	assert.True(t, v1)
	assert.False(t, v2)
}

func Test_deepNestedLogic(t *testing.T) {
	// Test with valueMap5
	v1 := ValidateExpression("((([1] AND [4]) OR ([2] AND [5])) AND [3]) OR [4]", valueMap5)
	// (((false OR false) AND (true OR false)) OR false) AND true
	v2 := ValidateExpression("((([1] OR [4]) AND ([2] OR [5])) OR [3]) AND [5]", valueMap6)
	// Test with valueMap1
	v3 := ValidateExpression("((([1] AND [2]) OR [3]) AND [1]) OR [2]", valueMap1)
	// Test with valueMap4
	v4 := ValidateExpression("((([1] AND [2]) OR [3]) AND [1]) OR [3]", valueMap4)

	assert.True(t, v1)
	assert.False(t, v2)
	assert.True(t, v3)
	assert.True(t, v4)
}
