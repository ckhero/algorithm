/**
 *@Description
 *@ClassName remove_duo
 *@Date 2021/4/6 下午5:12
 *@Author ckhero
 */

package main

import "fmt"

func main() {
	nums := []int{
		1,1,1,2,2,3,
	}
	fmt.Println(removeDuplicates(nums))
}

func removeDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	i := 1
	k := 1

	for j := 1; j < len(nums); j++ {
		// 当前位与前一位相等
		if nums[j] != nums[j - 1] {
			k = 1
			nums[i] = nums[j]
			i++
			continue
		}
		// 只出现过一次
		if k == 1 {
			k++
			nums[i] = nums[j]
			i++
			continue
		}
	}

	return i
}
