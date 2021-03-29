/**
 *@Description
 *@ClassName two_sum
 *@Date 2021/3/29 下午7:38
 *@Author ckhero
 */

package sum_two

func twoSum(nums []int, target int) []int {
	var (
		res = []int{}
		left = 0
		right = len(nums) - 1
	)
	for left < right {
		tmp := nums[left] + nums[right]
		if tmp == target {
			res = append(res, nums[left], nums[right])
			break
		}
		if tmp < target {
			left++
			continue
		}
		right--
	}
	return res
}
