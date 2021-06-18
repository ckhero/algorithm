/**
 *@Description
 *@ClassName missing_number
 *@Date 2021/6/18 上午10:15
 *@Author ckhero
 */

package _68_missing_number


func missingNumber(nums []int) int {
	res := len(nums)
	for k, num := range nums {
		res ^= k ^ num
	}
	return res
}