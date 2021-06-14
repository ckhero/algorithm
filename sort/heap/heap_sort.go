/**
 *@Description
 *@ClassName heap_sort
 *@Date 2021/6/1 下午4:43
 *@Author ckhero
 */

package heap

import (
	"fmt"
)

func HeapSort(data []int) {
	l := len(data)
	// 最后一个非页节点开始排序
	for i := 0; i < l; i++ {
		adjustHeap(data, l / 2 - 1, l - i)
		data[0], data[l - i - 1] = data[l - i - 1], data[0]
		fmt.Println(data)
	}
}

// 大顶堆
func adjustHeap(data []int, k, l int) {
	for i := k; i >= 0; i-- {
		curr := data[i]
		var left int
		var right int
		if (2 * i + 1) < l {
			left = data[2 * i + 1]
		}
		if (2 * i + 2) < l {
			right = data[2 * i + 2]
		}
		if left > curr && left >= right {
			data[2 * i + 1], data[i] = data[i], data[2 * i + 1]
		}
		if right > curr && right >= left {
			data[2 * i + 2], data[i] = data[i], data[2 * i + 2]
		}
	}
}