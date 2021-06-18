
/**
 *@Description
 *@ClassName missing_number_test.go
 *@Date 2021/6/18 上午10:21
 *@Author ckhero
 */

package _68_missing_number

import (
	"fmt"
	"testing"
)

func TestMissNums(t *testing.T) {
	nums := []int{9,6,4,2,3,5,7,0,1}
	fmt.Println(missingNumber(nums))
}