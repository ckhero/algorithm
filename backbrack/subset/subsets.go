/**
 *@Description
 *@ClassName subsets
 *@Date 2021/3/29 下午5:53
 *@Author ckhero
 */

package subset

func subsets(nums []int) [][]int {
	var helper func(nums, path []int, start int)
	var ans = [][]int{
		[]int{},
	}

	helper = func(nums, path []int, start int) {
		if len(nums) == len(path) || start == len(nums) {
			return
		}
		for i := start; i < len(nums); i++ {
			path = append(path, nums[i])
			tmp := make([]int, len(path))
			copy(tmp, path)
			ans = append(ans, tmp)
			helper(nums, path, i+1)
			path = path[:len(path)-1]
		}
	}
	helper(nums, []int{}, 0)
	return ans
}


