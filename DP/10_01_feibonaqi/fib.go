/**
 *@Description
 *@ClassName fib
 *@Date 2021/4/21 上午10:41
 *@Author ckhero
 */

package main

import (
	"fmt"
)

func main() {
	fmt.Println(fib(95))
}

func fib(n int) int {
	a := 0
	b := 1

	for i := 1; i <= n; i++ {
		b, a = (a % 1000000007 + b % 1000000007) %1000000007, b
	}

	return a
}
