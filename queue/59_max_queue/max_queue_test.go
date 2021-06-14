/**
 *@Description
 *@ClassName max_queue_test.go
 *@Date 2021/6/14 上午10:49
 *@Author ckhero
 */

package _9_max_queue

import (
	"fmt"
	"testing"
)

func TestConstructor(t *testing.T) {
	m := Constructor()
	m.Push_back(94)
	m.Push_back(16)
	m.Push_back(89)
	fmt.Println(m.Pop_front())
	m.Push_back(22)
	fmt.Println(m.Pop_front())

}