/**
 *@Description
 *@ClassName LRU
 *@Date 2021/5/26 上午10:53
 *@Author ckhero
 */

package LRU

import "container/list"

type LRUCache struct {
	ll    *list.List
	cache map[int]*list.Element
	cap   int
	len   int
}

type cacheEle struct {
	Key int
	Val int
}
func Constructor(capacity int) LRUCache {
	return LRUCache{
		ll:    list.New(),
		cache: map[int]*list.Element{},
		cap:   capacity,
		len:   0,
	}
}

func (this *LRUCache) Get(key int) int {
	if ele, ok := this.cache[key]; ok {
		this.ll.MoveToFront(ele)
		return ele.Value.(*cacheEle).Val
	}
	return -1
}

func (this *LRUCache) Put(key int, value int) {
	if ele, ok := this.cache[key]; ok {
		ele.Value.(*cacheEle).Val = value
		this.ll.MoveToFront(ele)
		return
	}
	ele := this.ll.PushFront(&cacheEle{
		Key: key,
		Val: value,
	})
	this.cache[key] = ele
	if this.len >= this.cap {
		delete(this.cache, this.ll.Back().Value.(*cacheEle).Key)
		this.ll.Remove(this.ll.Back())
	} else {
		this.len++
	}

}
