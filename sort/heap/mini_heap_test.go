/**
 *@Description
 *@ClassName mini_heap_test
 *@Date 2021/6/1 下午4:53
 *@Author ckhero
 */

package heap

import (
	"fmt"
	"testing"
)

func TestMiniHeap(t *testing.T) {
	data := []int{3,2,3,1,2,4,5,5,6}
	fmt.Println(MiniHeap(data, 4))
}
