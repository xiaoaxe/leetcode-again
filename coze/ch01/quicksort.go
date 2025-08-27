package ch01

import "fmt"

func Run_quicksort() {
	var nums = []int{5, 2, 1, 3, 4, 7}
	//nums = quicksort(nums)
	nums = quicksort2(nums)
	fmt.Printf("%+v\n", nums)
}

func quicksort(nums []int) []int {
	if len(nums) < 2 {
		return nums
	}
	var pilot = nums[0]
	var low, high []int
	for i := 1; i < len(nums); i++ {
		if nums[i] < pilot {
			low = append(low, nums[i])
		} else if nums[i] > pilot {
			high = append(high, nums[i])
		}
	}
	return append(append(quicksort(low), pilot), quicksort(high)...)
}

func quicksort2(nums []int) []int {
	if len(nums) < 2 {
		return nums
	}
	var pivot = nums[0]
	var pos = 1

	for j := len(nums) - 1; j > pos; j-- {
		for nums[j] < pivot && pos < j {
			swap(nums, pos, j)
			pos++
		}
	}
	if nums[pos] < pivot {
		swap(nums, pos, 0)
	}

	return append(append(quicksort2(nums[:pos]), nums[pos]), quicksort2(nums[pos+1:])...)
}

func swap(nums []int, i, j int) {
	nums[i], nums[j] = nums[j], nums[i]
}
