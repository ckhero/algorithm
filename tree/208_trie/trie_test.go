/**
 *@Description
 *@ClassName trie_test.go
 *@Date 2021/6/12 下午3:42
 *@Author ckhero
 */

package _08_trie

import (
	"fmt"
	"testing"
)

func TestConstructor(t *testing.T) {
	a := make([]*Trie, 26, 26)
	fmt.Println(a)
	trie := Constructor()
	trie.Insert("go")
	fmt.Println(trie.Search("go"))
	fmt.Println(trie.Search("golang"))
	fmt.Println(trie.Freq("go"))
	fmt.Println(trie.Freq("g"))
	trie.Insert("golang")
	fmt.Println(trie.Search("go"))
	fmt.Println(trie.Search("golang"))
	fmt.Println(trie.Freq("go"))
	fmt.Println(trie.Freq("g"))
}