package datastructures

import (
	"strings"
)

const terminateChar = "*"

/*
Trie returns a pointer to a trie type. The `root` node is populated
with a pointer to a trieNode type.
*/
func Trie() *trie {
	return &trie{
		root: &trieNode{
			children: make(map[string]*trieNode),
		},
	}
}

/*
Root returns a pointer to the root node of the trie.
*/
func (t *trie) Root() *trieNode {
	return t.root
}

/*
Search attempts to find the given `word` among all the nodes within the trie.

If the word can be found, it returns a pointer to the node that contains the
terminating character, `*`.

This is for the purposes of implementing the Autocomplete function.
*/
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

/*
Insert breaks a `word` up separate letters, traversing the trie and inserting
child nodes where a given letter in `word` does not already exist in the trie.

Each word is terminated by the `*` character.
*/
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

/*
Autocomplete takes a `prefix` string and lists words that begin with it.
It returns `nil` if no words can be found.
*/
func (t *trie) Autocomplete(prefix string) []string {
	currentNode := t.Search(prefix)
	if currentNode == nil {
		return nil
	}

	return t.list(currentNode, prefix, []string{})
}

/*
list traverses the tree and attempts to find words that can be terminated with
the given `word` argument.
*/
func (t *trie) list(node *trieNode, word string, words []string) []string {
	currentNode := node

	for letter, child := range currentNode.children {
		if letter == terminateChar {
			words = append(words, word)
			continue
		}

		words = t.list(child, word+letter, words)
	}

	return words
}

/*
Traverse collects all the letters of the child nodes within a trie
and returns a slice of letters.
*/
func (t *trie) Traverse(node *trieNode, letters []string) []string {
	for letter, child := range node.children {
		letters = t.Traverse(
			child,
			append(letters, letter),
		)
	}

	return letters
}
