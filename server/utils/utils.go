package utils

import "time"

const defaultTImeLayout = "2006:01:02 15:04:05"

func RemoveDuplicates(nums []int) []int {
	resultMap := make(map[int]bool)
	var result []int

	for _, num := range nums {
		if !resultMap[num] {
			result = append(result, num)
			resultMap[num] = true
		}
	}

	return result
}

func FormatTime(t time.Time, layout string) string {
	if layout == "" {
		return t.Format(defaultTImeLayout)
	}
	return t.Format(layout)
}
