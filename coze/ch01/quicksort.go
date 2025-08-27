package ch01

import "fmt"

func Run_quicksort() {
	//var nums = []int{5, 2, 1, 3, 4, 7}
	//var nums = []int{5, 2, 1, 6, 4, 7}
	var nums = []int{9, 2, 1, 6, 4, 7}
	//nums = quicksort(nums)
	quicksort2(nums)
	fmt.Printf("%+v\n", nums)
}

func quicksort(nums []int) []int {
	if len(nums) < 2 {
		return nums
	}
	var pivot = nums[0]
	var low, high []int
	for i := 1; i < len(nums); i++ {
		if nums[i] < pivot {
			low = append(low, nums[i])
		} else if nums[i] > pivot {
			high = append(high, nums[i])
		}
	}
	return append(append(quicksort(low), pivot), quicksort(high)...)
}

func quicksort2(nums []int) {
	if len(nums) < 2 {
		return
	}
	var pivot = nums[0]
	var pos = 1
	for j := len(nums) - 1; j > pos; j-- {
		for nums[j] < pivot && pos < j {
			swap(nums, pos, j)
			pos++
		}
		//如果j位置的值更大的话，需要替换前面的一位
		if pos == j && nums[j] > pivot {
			pos--
		}
	}

	//把pivot移动到合适的位置
	if nums[pos] < pivot {
		swap(nums, pos, 0)
	} else {
		pos = 0
	}

	quicksort2(nums[:pos])
	quicksort2(nums[pos+1:])
}

func swap(nums []int, i, j int) {
	nums[i], nums[j] = nums[j], nums[i]
}
