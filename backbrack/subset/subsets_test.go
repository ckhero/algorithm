/**
 *@Description
 *@ClassName subsets_test
 *@Date 2021/3/29 下午6:04
 *@Author ckhero
 */

package subset

import (
	"fmt"
	"testing"
)

func TestSubset(t *testing.T) {
	nums := []int{1, 2, 3}
	ans := subsets(nums)
	fmt.Println(ans)
}

