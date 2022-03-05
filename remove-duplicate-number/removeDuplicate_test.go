package removeduplicatenumber

import (
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

func RmDupSolution1(nums []int) []int {
	var result []int
	sort.Ints(nums)
	for i := 0; i < len(nums); i++ {
		isDuplicate := false
		for j := i; j < len(nums); j++ {
			if j == i {
				continue
			} else if nums[j] == nums[i] {
				isDuplicate = true
				break
			}
		}
		if !isDuplicate {
			result = append(result, nums[i])
		}
	}
	return result
}

func RmDupSolution2(nums []int) []int {
	var result []int
	sort.Ints(nums)
	for i := 1; i < len(nums); i++ {
		if nums[i] != nums[i-1] {
			result = append(result, nums[i-1])
		}
	}
	result = append(result, nums[len(nums)-1])
	return result
}

func Test_RemoveDuplicate(t *testing.T) {
	t.Run("Test_RemoveDupSolution1", func(t *testing.T) {
		assert.Equal(t, []int{1, 2, 3, 4, 5}, RmDupSolution1([]int{1, 1, 2, 3, 4, 4, 5}))
		assert.Equal(t, []int{1, 2, 3, 4, 5}, RmDupSolution1([]int{1, 3, 3, 2, 4, 4, 5}))
		assert.Equal(t, []int{1, 2, 3, 4, 5, 6}, RmDupSolution1([]int{1, 1, 2, 2, 2, 3, 4, 4, 5, 6, 6, 6}))
		assert.Equal(t, []int{1, 2, 3, 4, 5, 6}, RmDupSolution1([]int{1, 3, 3, 2, 4, 4, 5, 5, 6, 6, 6}))
		assert.Equal(t, []int{1, 2}, RmDupSolution1([]int{1, 1, 1, 2, 2, 1, 1, 2, 1}))
		// PASS ok
	})
	t.Run("Test_RemoveDupSolution2", func(t *testing.T) {
		assert.Equal(t, []int{1, 2, 3, 4, 5}, RmDupSolution2([]int{1, 1, 2, 3, 4, 4, 5}))
		assert.Equal(t, []int{1, 2, 3, 4, 5, 6}, RmDupSolution2([]int{1, 1, 2, 2, 2, 3, 4, 4, 5, 6, 6, 6}))
		assert.Equal(t, []int{1, 2, 3, 4, 5}, RmDupSolution2([]int{1, 3, 3, 2, 4, 4, 5}))
		assert.Equal(t, []int{1, 2, 3, 4, 5, 6}, RmDupSolution2([]int{1, 3, 3, 2, 4, 4, 5, 5, 6, 6, 6}))
		assert.Equal(t, []int{1, 2}, RmDupSolution2([]int{1, 1, 1, 2, 2, 1, 1, 2, 1}))
		// PASS ok
	})
}
