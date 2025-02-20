package main

func singleNonDuplicate(nums []int) int {

	resChan := make(chan int, 1)

	if len(nums) == 1 {
		return nums[0]
	}

	go func([]int) {
		lastEl := nums[0]
		counter := 0
		for _, el := range nums {
			if lastEl == el {
				counter++
				continue
			} else {
				if counter < 2 {
					resChan <- lastEl
					break
				}
				counter = 1
				lastEl = el
			}
		}
	}(nums)

	go func([]int) {
		lastEl := nums[len(nums)-1]
		counter := 0
		for i := len(nums) - 1; i >= 0; i-- {
			if lastEl == nums[i] {
				counter++
				continue
			} else {
				if counter < 2 {
					resChan <- lastEl
					break
				}
				counter = 1
				lastEl = nums[i]
			}
		}
	}(nums)

	return <-resChan
}
