/**
 *@Description
 *@ClassName merge_sort
 *@Date 2021/6/5 下午5:13
 *@Author ckhero
 */

package merge_sort

import "fmt"

func MergeSort(data []int) []int {
	if len(data) == 1 {
		return data
	}
	mid := len(data) / 2
	left := MergeSort(data[0:mid])
	right := MergeSort(data[mid:])
	return Merge(left, right)
}

func Merge(left, right []int) []int {
	res := []int{}
	fmt.Println(cap(res))
	lIndex := 0
	rIndex := 0
	for lIndex < len(left) && rIndex < len(right) {
		if left[lIndex] <= right[rIndex] {
			res = append(res, left[lIndex])
			lIndex++
			continue
		}
		res = append(res, right[rIndex])
		rIndex++
	}
	res = append(res, left[lIndex:]...)
	res = append(res, right[rIndex:]...)
	return res
}
