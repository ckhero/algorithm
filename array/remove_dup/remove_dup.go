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
	slow := 1
	k := 1

	for fast := 1; fast < len(nums); fast++ {
		// 当前位与前一位相等
		if nums[fast] != nums[fast - 1] {
			k = 1
			nums[slow] = nums[fast]
			slow++
			continue
		}
		// 只出现过一次
		if k == 1 {
			k++
			nums[slow] = nums[fast]
			slow++
			continue
		}
	}

	return i
}
