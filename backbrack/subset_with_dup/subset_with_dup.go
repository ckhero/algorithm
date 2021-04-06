/**
 *@Description
 *@ClassName subset_with_dup
 *@Date 2021/3/31 下午4:31
 *@Author ckhero
 */

package main

import (
	"fmt"
	"sort"
)

func main() {
	nums := []int{4,4,4,1,4}
	fmt.Println(subsetsWithDup(nums))
}

func subsetsWithDup(nums []int) [][]int {
	sort.Ints(nums)
	var (
		ans = [][]int{[]int{}}
	)
	var helper func(nums []int, path []int, start int)
	helper = func(nums []int, path []int, start int) {
		if len(path) == len(nums) || start >= len(nums) {
			return
		}
		for i := start; i < len(nums); i++ {
			if i > start && nums[i] == nums[i-1] {
				continue
			}
			path = append(path, nums[i])
			tmp := make([]int, len(path))
			copy(tmp, path)
			ans = append(ans, tmp)

			helper(nums, path, i+1)
			path = path[0 : len(path)-1]
		}
	}
	helper(nums, []int{}, 0)
	return ans
}
