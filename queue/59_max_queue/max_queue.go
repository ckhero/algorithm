/**
 *@Description
 *@ClassName max_queue
 *@Date 2021/6/14 上午10:43
 *@Author ckhero
 */

package _9_max_queue

type MaxQueue struct {
	queue []int
	maxQueue []int
}


func Constructor() MaxQueue {
	return MaxQueue{
		queue:    []int{},
		maxQueue: []int{},
	}
}


func (this *MaxQueue) Max_value() int {
	if len(this.maxQueue) == 0 {
		return -1
	}
	return this.maxQueue[0]
}


func (this *MaxQueue) Push_back(value int)  {
	this.queue = append(this.queue, value)
	for i := len(this.maxQueue) - 1; i >= 0 ; i-- {
		if this.maxQueue[i] <= value {
			this.maxQueue = this.maxQueue[:len(this.maxQueue) - 1]
		}
	}
	this.maxQueue = append(this.maxQueue, value)
}


func (this *MaxQueue) Pop_front() int {
	if len(this.queue) == 0 {
		return -1
	}
	val := this.queue[0]
	if val == this.maxQueue[0] {
		this.maxQueue = this.maxQueue[1:]
	}
	this.queue = this.queue[1:]
	return val
}
