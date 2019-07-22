package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBirthday(t *testing.T) {
	t.Run("when slice length is smaller than 1", func(t *testing.T) {
		nums := []int32{0}
		actual := Birthday(nums, 1, 1)
		expected := int32(0)
		assert.Equal(t, expected, actual)
	})

	t.Run("when slice contains 1 element with correct m and d", func(t *testing.T) {
		nums := []int32{1}
		actual := Birthday(nums, 1, 1)
		expected := int32(1)
		assert.Equal(t, expected, actual)
	})

	t.Run("when slice contains 1 element with incorrect d", func(t *testing.T) {
		nums := []int32{1}
		actual := Birthday(nums, 2, 1)
		expected := int32(0)
		assert.Equal(t, expected, actual)
	})

	t.Run("when slice contains 1 element with incorrect m", func(t *testing.T) {
		nums := []int32{1}
		actual := Birthday(nums, 1, 2)
		expected := int32(0)
		assert.Equal(t, expected, actual)
	})

	t.Run("when m is out of range", func(t *testing.T) {
		nums := []int32{1, 2, 3}
		actual := Birthday(nums, 3, 13)
		expected := int32(0)
		assert.Equal(t, expected, actual)
	})

	t.Run("when d is out of range", func(t *testing.T) {
		nums := []int32{1, 2, 3}
		actual := Birthday(nums, 32, 2)
		expected := int32(0)
		assert.Equal(t, expected, actual)
	})

	t.Run("when s length is out of range", func(t *testing.T) {
		nums := make([]int32, 101)
		actual := Birthday(nums, 3, 2)
		expected := int32(0)
		assert.Equal(t, expected, actual)
	})

	t.Run("when s, m and d are all correct", func(t *testing.T) {
		nums := []int32{1, 2, 1, 3, 2}
		actual := Birthday(nums, 3, 2)
		expected := int32(2)
		assert.Equal(t, expected, actual)
	})
}
