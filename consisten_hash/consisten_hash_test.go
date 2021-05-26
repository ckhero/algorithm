/**
 *@Description
 *@ClassName consisten_hash_test
 *@Date 2021/5/26 上午10:05
 *@Author ckhero
 */

package consisten_hash

import (
	"fmt"
	"hash/crc32"
	"strconv"
	"testing"
)

func TestNewConsistenHash(t *testing.T) {
	hash := NewConsistenHash(1, crc32.ChecksumIEEE)
	hash.AddMachine("123")
	hash.AddMachine("2")
	hash.AddMachine("123123123123123")
	fmt.Println(hash.GetMachine("test1"))
	fmt.Println(hash.GetMachine("aaa发送到发发士大夫"))
	fmt.Println(hash.GetMachine("test1"))
	hash.AddMachine("test051")
	hash.AddMachine("fadsfadfadfssadf")
	hash.AddMachine("aaaaaaaaaa")
	hash.AddMachine("22222")
	hash.AddMachine("1")
	fmt.Println(hash.GetMachine("test1"))
	for i := 0; i < 100; i ++ {
		fmt.Println(i, hash.GetMachine(strconv.Itoa(i)))

	}
}
