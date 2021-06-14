/**
 *@Description
 *@ClassName trie_with_doc_id
 *@Date 2021/6/12 下午3:53
 *@Author ckhero
 */

package _08_trie

type TriceWithDocId struct {
	chidlNodes []*Trie
	docMap map[int]struct{}
}
