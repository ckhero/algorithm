/**
 *@Description
 *@ClassName hamming_weight_test
 *@Date 2021/5/31 下午3:13
 *@Author ckhero
 */

package _5_hamming_weight

import (
	"fmt"
	"strconv"
	"testing"
)

func Test(t *testing.T) {
	a := uint64(200)
	fmt.Println(strconv.FormatUint(a, 16))
	fmt.Println(hammingWeightV1(9))
}
