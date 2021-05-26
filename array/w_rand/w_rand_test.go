/**
 *@Description
 *@ClassName w_rand_test
 *@Date 2021/5/24 下午3:51
 *@Author ckhero
 */

package w_rand

import (
	"fmt"
	"testing"
)

func TestRand(t *testing.T) {
	w := []int{1,2,3}
	s := Constructor(w)
	fmt.Println(s.PickIndex())
	fmt.Println(s.PickIndex())
	fmt.Println(s.PickIndex())
}
