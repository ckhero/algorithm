/**
 *@Description
 *@ClassName insert_sort
 *@Date 2021/6/6 下午3:45
 *@Author ckhero
 */

package insert_sort

func InsertSort(data []int) []int {
	for i := 1; i < len(data) - 1; i++ {
		for j := i + 1; j > 0; j-- {
			if data[j] < data[j - 1] {
				data[j], data[j - 1] = data[j - 1], data[j]
			}
		}
	}
	return data
}
