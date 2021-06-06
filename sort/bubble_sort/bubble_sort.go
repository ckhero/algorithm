/**
 *@Description
 *@ClassName bubble_sort
 *@Date 2021/6/6 下午3:27
 *@Author ckhero
 */

package bubble_sort

func BubbleSort(data []int) []int{
	isSorted := true
	for i := 0; i < len(data); i++ {
		for j := 1; j < len(data) - i; j++ {
			if data[j] < data[j - 1] {
				data[j], data[j - 1] = data[j-1], data[j]
				isSorted = false
			}
		}
		if isSorted {
			// 最优情况下 O(n)
			break
		}

	}
	return data
}
