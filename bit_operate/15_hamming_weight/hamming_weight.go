/**
 *@Description
 *@ClassName hamming_weight
 *@Date 2021/5/31 下午3:10
 *@Author ckhero
 */

package _5_hamming_weight

func hammingWeight(num uint32) int {
	var count int
	for num > 0 {
		count += int(num & 1)
		num = num >> 1
	}
	return count
}


func hammingWeightV1(num uint32) int {
	var count int
	for num > 0 {
		count++
		num = num & (num - 1)
	}
	return count
}