/**
 *@Description
 *@ClassName quick_sort
 *@Date 2021/6/5 下午3:26
 *@Author ckhero
 */

package quick

func QuickSort(data []int, low, high int) {
	if low >= high {
		return
	}
	var (
		base = data[low]
		left = low
		right = high
	)

	for left != right {
		for left < right && data[right] > base {
			right--
		}

		for left < right && data[left] <= base {
			left++
		}

		if left < right {
			data[left], data[right] = data[right], data[left]
		}
	}
	data[low], data[left] = data[left], data[low]
	QuickSort(data, low, left - 1)
	QuickSort(data, left + 1, high)
	return
}
