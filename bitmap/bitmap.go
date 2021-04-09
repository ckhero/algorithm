/**
 *@Description
 *@ClassName bitmap
 *@Date 2021/4/9 下午2:02
 *@Author ckhero
 */

package main

import "fmt"

func main() {
	b := NewBitMap()
	b.Mark(100)
	fmt.Println(b.IsMark(100))
	fmt.Println(b.IsMark(11))
	b.Mark(11)
	fmt.Println(b.IsMark(11))
	fmt.Println(b.MakLen())
}

type BitMap struct {
	body []uint8
	markLen int
}

func NewBitMap() *BitMap {
	return &BitMap{
		body:    []uint8{},
		markLen: 0,
	}
}
func(b *BitMap) Mark(index int)  {
	pos := index /  8
	offset := index % 8
	if offset != 0 {
		pos ++
	} else {
		offset = 8
	}
	if pos > len(b.body) {
		appendBody := make([]uint8, pos - len(b.body))
		b.body = append(b.body, appendBody...)
	}
	tmp := uint8(128)
	for offset > 1 {
		tmp >>= 1
		offset--
	}
	if (b.body[pos - 1] & tmp) == tmp {
		return
	}
	b.body[pos - 1] ^= tmp
	b.markLen++

}


func(b *BitMap) IsMark(index int) bool {
	pos := index /  8
	offset := index % 8
	if offset != 0 {
		pos ++
	} else {
		offset = 8
	}
	if pos > len(b.body) {
		appendBody := make([]uint8, pos - len(b.body))
		b.body = append(b.body, appendBody...)
	}
	tmp := uint8(128)
	for offset > 1 {
		tmp >>= 1
		offset--
	}
	 return (b.body[pos - 1] & tmp) == tmp
}

func (b *BitMap) MakLen() int {
	return b.markLen
}