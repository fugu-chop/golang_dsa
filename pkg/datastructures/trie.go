package datastructures

import "strings"

/*
Trie returns a pointer to a trie type. The `root` node is populated
with a pointer to a trieNode type with the key of `letter`.
*/
func Trie(letter string) *trie {
	return &trie{
		root: &trieNode{
			children: make(map[string]*trieNode),
		},
	}
}

func (t *trie) Search(word string) *trieNode {
	currentNode := t.root
	strSlice := strings.Split(word, "")
	for _, letter := range strSlice {
		if currentNode.Get(letter) == nil {
			return nil
		}
		currentNode = currentNode.Get(letter)
	}
	return currentNode
}

func (t *trie) Insert(word string) {
	currentNode := t.root
	strSlice := strings.Split(word, "")
	for _, letter := range strSlice {
		if currentNode.Get(letter) != nil {
			currentNode = currentNode.Get(letter)
			continue
		}

		currentNode.Set(letter)
		currentNode = currentNode.Get(letter)
	}

	currentNode.Set("*")
}
