/**
 *@Description
 *@ClassName w_rand
 *@Date 2021/5/24 下午3:42
 *@Author ckhero
 */

package w_rand

import (
	"math/rand"
)

type Solution struct {
	wSums     []int
	total int
}

func Constructor(w []int) Solution {
	var total int
	wSums := make([]int, len(w))
	for key, val := range w {
		total += val
		wSums[key] = total
	}
	return Solution{wSums, total}
}

func (this *Solution) PickIndex() int {
	target := rand.Intn(this.total)
	left := 0
	right := len(this.wSums)
	for left < right {
		mid := (left + right) / 2
		if target >= this.wSums[mid] {
			left = mid + 1
		} else if target < this.wSums[mid] {
			right = mid
		}
	}
	return left
}
