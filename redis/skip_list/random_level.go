/**
 *@Description
 *@ClassName random_level
 *@Date 2021/4/8 上午9:16
 *@Author ckhero
 */

package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	//rand.Seed(time.Now().UnixNano())
	res := []int{0, 0, 0, 0, 0}
	for i := 0; i < 100000; i++ {
		index := randomLevel()
		res[index]++
	}

	fmt.Println(res)
}


/**
返回1的概率为 1/2
返回2的概率为 1/4
返回3的概率为 1/8

返回 2的时候建立 一级索引
返回 3的时候建立 二级索引
以此类推
 */
func randomLevel() int {

	maxLevel := 3
	level := 1
	rand.Seed(time.Now().UnixNano())
	for rand.Intn(2) == 1 && level <= maxLevel {
		level++
	}
	return level
}
