package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBreakingRecords(t *testing.T) {
	t.Run("when less than two scores", func(t *testing.T) {
		scores := []int32{1}

		actual := BreakingRecords(scores)
		expected := []int32{0, 0}

		assert.Equal(t, expected, actual)
	})

	t.Run("when two scores with the same values", func(t *testing.T) {
		scores := []int32{2, 2}

		actual := BreakingRecords(scores)
		expected := []int32{0, 0}

		assert.Equal(t, expected, actual)
	})

	t.Run("when two scores with different values", func(t *testing.T) {
		scores := []int32{2, 1}

		actual := BreakingRecords(scores)
		expected := []int32{0, 1}

		assert.Equal(t, expected, actual)
	})

	t.Run("when more scores with the same values", func(t *testing.T) {
		scores := []int32{2, 2, 2}

		actual := BreakingRecords(scores)
		expected := []int32{0, 0}

		assert.Equal(t, expected, actual)
	})

	t.Run("when more scores with different values", func(t *testing.T) {
		scores := []int32{2, 1, 3}

		actual := BreakingRecords(scores)
		expected := []int32{1, 1}

		assert.Equal(t, expected, actual)
	})
}
