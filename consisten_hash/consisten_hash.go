/**
 *@Description
 *@ClassName consisten_hash
 *@Date 2021/5/26 上午9:50
 *@Author ckhero
 */

package consisten_hash

import (
	"sort"
	"strconv"
)

type Hash func(data []byte) uint32

type consistenHash struct {
	replicas   int
	hash       Hash
	circle     []int
	machineMap map[int]string
}

func NewConsistenHash(replicas int, hash Hash) *consistenHash {
	return &consistenHash{
		replicas:   replicas,
		hash:       hash,
		circle:     []int{},
		machineMap: map[int]string{},
	}
}

func (c *consistenHash) AddMachine(machine string) {
	for i := 0; i < c.replicas; i++ {
		hashVal := int(c.hash([]byte(machine+strconv.Itoa(i) )))
		c.circle = append(c.circle, hashVal)
		c.machineMap[hashVal] = machine
	}
	sort.Ints(c.circle)
}

func (c *consistenHash) DelMachine(machine string) {
	for i := 0; i < c.replicas; i++ {
		hashVal := int(c.hash([]byte(machine+strconv.Itoa(i) )))
		c.circle = append(c.circle, hashVal)
		c.machineMap[hashVal] = machine
	}
	sort.Ints(c.circle)
}

func (c *consistenHash) GetMachine(key string) string {
	hashVal := int(c.hash([]byte(key)))
	//fmt.Println(c.circle)
	//fmt.Println(c.machineMap)
	//fmt.Println(hashVal)
	var currKey int
	var preVal int
	for k, val := range c.circle {
		if val >= hashVal && preVal < hashVal {
			currKey = k
			preVal = val
		}
	}
	return c.machineMap[c.circle[currKey%len(c.circle)]]
}
