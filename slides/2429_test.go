package slides

import (
	"math/bits"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMinimizeXor(t *testing.T) {
	num1 := 3
	num2 := 5
	res := minimizeXor(num1, num2)
	assert.Equal(t, 3, res)
	num1 = 25
	num2 = 72
	res = minimizeXor(num1, num2)
	assert.Equal(t, 24, res)
}

// minimizeXor()
func minimizeXor(num1 int, num2 int) int {
	c1 := bits.OnesCount(uint(num1))
	c2 := bits.OnesCount(uint(num2))

	for ; c2 < c1; c2++ {
		num1 &= num1 - 1 // 末尾1变0
	}

	for ; c2 > c1; c1++ {
		num1 |= num1 + 1 // 末尾0变1
	}
	return num1
}
