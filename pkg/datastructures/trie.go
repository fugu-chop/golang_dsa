package datastructures

import (
	"fmt"
	"strings"
)

const terminateChar = "*"

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

	currentNode.Set(terminateChar)
}

func (t *trie) Autocomplete(prefix string) []string {
	currentNode := t.Search(prefix)
	if currentNode == nil {
		return nil
	}

	return t.list(currentNode, prefix, []string{})
}

func (t *trie) list(node *trieNode, word string, words []string) []string {
	currentNode := node

	for letter, child := range currentNode.children {
		if letter == terminateChar {
			words = append(words, word)
			continue
		}

		t.list(child, word+letter, words)
	}

	return words
}

func (t *trie) Traverse(node *trieNode) {
	if node.children == nil {
		return
	}

	for letter, child := range node.children {
		fmt.Println(letter)
		t.Traverse(child)
	}
}
