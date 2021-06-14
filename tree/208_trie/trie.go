/**
 *@Description
 *@ClassName trie
 *@Date 2021/6/12 下午3:29
 *@Author ckhero
 */

package _08_trie

type Trie struct {
	childNodes []*Trie
	isEnd bool
	freq int
}


/** Initialize your data structure here. */
func Constructor() Trie {
	return Trie{
		childNodes: make([]*Trie, 26, 26),
	}
}


/** Inserts a word into the trie. */
func (this *Trie) Insert(word string)  {
	currNode := this
	for _, ch := range word {
		ch -= 'a'
		if currNode.childNodes[ch] == nil {
			currNode.childNodes[ch] = &Trie{
				childNodes: make([]*Trie, 26, 26),
			}
		}
		currNode.childNodes[ch].freq++
		currNode = currNode.childNodes[ch]
	}
	currNode.isEnd = true
}


/** Returns if the word is in the trie. */
func (this *Trie) Search(word string) bool {
	node := this.StartPrefix(word)
	return node != nil && node.isEnd
}

func(this *Trie) StartPrefix(prefix string) *Trie {
	currNode := this
	for _, ch := range prefix {
		ch -= 'a'
		if currNode.childNodes[ch] == nil {
			return nil
		}
		currNode = currNode.childNodes[ch]
	}
	return currNode
}

/** Returns if there is any word in the trie that starts with the given prefix. */
func (this *Trie) StartsWith(prefix string) bool {
	node := this.StartPrefix(prefix)
	return node != nil
}

func (this *Trie) Freq(prefix string) int {
	node := this.StartPrefix(prefix)
	if node == nil {
		return 0
	}
	return node.freq
}
/**
 * Your Trie object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(word);
 * param_2 := obj.Search(word);
 * param_3 := obj.StartsWith(prefix);
 */
