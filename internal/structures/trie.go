/*
Incomplete implementation. Please ignore!
*/

package structures

// fmt temporary
import (
	"fmt"
)

type TrieNode struct {
	val      rune
	children map[rune]*TrieNode
	final    bool
}

func (tnode *TrieNode) AddChild(char rune) {
	if tnode.children == nil {
		tnode.children = make(map[rune]*TrieNode)

	}

	// addr to new trienode
	newchild := &TrieNode{
		val:      char,
		children: make(map[rune]*TrieNode),
	}
	tnode.children[char] = newchild
}

func (tnode *TrieNode) RemoveDirectChild(char rune) {
	if tnode.children == nil {
		return
	}

	ckey, cnode := tnode.children[char]

	if !cnode {
		return
	}

	if len(ckey.children) == 0 && ckey.final {
		delete(tnode.children, char)
		return
	}
}

func (tnode *TrieNode) RemoveString(str string) {
	toslice := []rune(str)
	for i := len(toslice) - 1; i >= 0; i-- {
		fmt.Printf("%c", toslice[i])
	}

}

// func (root *TrieNode) getAllStrings() []string {
// 	var results []string
// 	for key, value := range root.children {

// 	}
// 	return
// }
