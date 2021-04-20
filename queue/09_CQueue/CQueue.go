/**
 *@Description
 *@ClassName CQueue
 *@Date 2021/4/20 下午8:07
 *@Author ckhero
 */

package _9_CQueue

type CQueue struct {
	queue []int
}

func Constructor() CQueue {
	return CQueue{queue: []int{}}
}

func (this *CQueue) AppendTail(value int) {
	this.queue = append(this.queue, value)
}

func (this *CQueue) DeleteHead() int {
	if len(this.queue) == 0 {
		return -1
	}
	res := this.queue[0]
	this.queue = this.queue[1:]
	return res
}
