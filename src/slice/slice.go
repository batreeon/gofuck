package slice

import (
	"slices"
	"sort"
)

// DeleteIndexes 使用下标，从s中删除多个元素
func DeleteIndexes[T any](s []T, indexes []int) []T {
	sort.Sort(sort.Reverse(sort.IntSlice(indexes)))
	for _, index := range indexes {
		if index >= len(s) || index < 0 {
			continue
		}
		s = slices.Delete(s, index, index+1)
	}
	return s
}
