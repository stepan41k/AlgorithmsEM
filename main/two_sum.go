package main

// 1. Two Sum

func twoSum(nums []int, target int) []int {

    numsMap := make(map[int]int)

    for i, num := range nums {
		diff := target - num

        if idx, found := numsMap[diff]; found {
            return []int{i, idx}
        }
        
    	numsMap[num] = i
	}


    return nil
}