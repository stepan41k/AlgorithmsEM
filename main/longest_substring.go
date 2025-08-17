package main

//3. Longest Substring Without Repeating Characters

func lengthOfLongestSubstring(s string) int {
    maxCount := 0
	left := 0
	idxes := make(map[byte]int)

    for right := 0; right < len(s); right++ {
        if val, exists := idxes[s[right]]; exists {
            left = max(left, val+1)
        }
		idxes[s[right]] = right
        maxCount = max(maxCount, right-left+1)
	}

    return maxCount
}