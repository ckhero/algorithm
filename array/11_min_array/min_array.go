/**
 *@Description
 *@ClassName min_array
 *@Date 2021/4/22 上午10:56
 *@Author ckhero
 */

package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main()  {
	resp, err := http.Get(`http://merchant-interface.test.zk022.cn/api/v2/order/batchCancelOrder?appId=9740ec5fb15f`)

	fmt.Println("http.Get: ", err)

	if err == nil {
		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)

		fmt.Println(`cancelOrderUrl resp: `, string(body))
		if err != nil {
			fmt.Println(`orderServer error:`, err)
			return
		}
	}
	//fmt.Println(minArray([]int{2,2,2,0,1}))
	//fmt.Println(minArray([]int{1, 3, 5}))
}
func minArray(numbers []int) int {
	left := 0
	right := len(numbers) - 1
	min := numbers[0]
	for left < right  {
		mid := (left + right) / 2
		if numbers[mid] >= numbers[left] {
			left = mid + 1
			continue
		}
		right = mid - 1
	}
	return min
}

// O(n)
func minArray1(numbers []int) int {
	if len(numbers) == 0 {
		return -1
	}

	min := numbers[0]
	for i := 1; i < len(numbers); i++  {
		if numbers[i]  < min {
			return numbers[i]
		}
	}
	return min
}