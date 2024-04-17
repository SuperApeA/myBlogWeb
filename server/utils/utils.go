package utils

import (
	"crypto/md5"
	"fmt"
	"time"
)

const (
	DefaultTImeLayout = "2006:01:02 15:04:05"
	Md5ExtraStr       = "AajBlog"
)

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
		return t.Format(DefaultTImeLayout)
	}
	return t.Format(layout)
}

func Md5Crypt(str string) string {
	return fmt.Sprintf("%x", md5.Sum([]byte(str+Md5ExtraStr)))
}
