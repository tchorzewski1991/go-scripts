package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestKangaroo(t *testing.T) {

	var x1, v1, x2, v2 int32

	t.Run("Returns positive result for correct set of values",
		func(t *testing.T) {
			x1, v1 = 0, 3
			x2, v2 = 4, 2

			actual := Kangaroo(x1, v1, x2, v2)
			expected := "YES"

			assert.Equal(t, expected, actual)
		})

	t.Run("Returns negative result for incorrect set of values",
		func(t *testing.T) {
			x1, v1 = 0, 2
			x2, v2 = 5, 3

			actual := Kangaroo(x1, v1, x2, v2)
			expected := "NO"

			assert.Equal(t, expected, actual)
		})

	t.Run("Returns negative result when any of value is out of limit",
		func(t *testing.T) {
			x1, v1 = 4523, 8092
			x2, v2 = 9419, 8076

			actual := Kangaroo(x1, v1, x2, v2)
			expected := "NO"

			assert.Equal(t, expected, actual)
		})
}
