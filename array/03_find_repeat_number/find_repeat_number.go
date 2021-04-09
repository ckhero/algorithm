/**
 *@Description
 *@ClassName find_repeat_number
 *@Date 2021/4/8 下午7:18
 *@Author ckhero
 */

package _3_find_repeat_number

func findRepeatNumber(nums []int) int {
	for i := 0;i < len(nums); i++ {
		if i != nums[i] {
			if nums[i] == nums[nums[i]] {
				return nums[i]
			} else {
				nums[i], nums[nums[i]] = nums[nums[i]], nums[i]
			}
		}
	}
	return -1
}
