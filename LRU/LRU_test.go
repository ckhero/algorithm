/**
 *@Description
 *@ClassName LRU_test
 *@Date 2021/5/26 上午11:06
 *@Author ckhero
 */

package LRU

import (
	"fmt"
	"testing"
)

func TestConstructor(t *testing.T) {
	c := Constructor(2)
	fmt.Println(c.Get(1))
	c.Put(1,1)
	fmt.Println(c.Get(1))

	c.Put(2,2)
	c.Put(3,3)
	fmt.Println(c.Get(1))

	c.Put(4,4)
	fmt.Println(c.Get(1))
	fmt.Println(c.Get(2))
	fmt.Println(c.Get(3))

	c.Put(5,5)
	fmt.Println(c.Get(3))
	fmt.Println(c.Get(4))


}
