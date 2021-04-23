/**
 *@Description
 *@ClassName num_ways
 *@Date 2021/4/22 上午10:47
 *@Author ckhero
 */

package main

import "fmt"

func main() {
	fmt.Println(numWays(0))
	fmt.Println(numWays(1))
	fmt.Println(numWays(2))
}

func numWays(n int) int {
	a := 0
	b := 1
	for i := 0; i <= n; i ++ {
		b, a = (a % 1000000007 + b % 1000000007) % 1000000007, b
	}
	return a
}