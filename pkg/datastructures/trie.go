package datastructures

import (
	"strings"
)

const terminateChar = "*"

/*
trie is an implementation of a trie (retrieval) data structure.
It holds a `root`, which is a pointer to a trieNode.
*/
type trie struct {
	root *trieNode
}

/*
a trieNode is a node used within the implementation of a Trie.
It contains a map whose keys are letters and values are pointers to
other trieNodes.
*/
type trieNode struct {
	children map[string]*trieNode
}

/*
Get attempts to fetch a pointer to a child node that has a child with a value of `letter`.
*/
func (t *trieNode) get(letter string) *trieNode {
	return t.children[letter]
}

/*
Set creates a child node for the current node that has a `letter` value.
If a child node already exists with a `letter` value, `Set` is a no-op.
*/
func (t *trieNode) set(letter string) {
	// avoid clobbering existing relationships
	if _, ok := t.children[letter]; ok {
		return
	}

	t.children[letter] = &trieNode{
		children: make(map[string]*trieNode),
	}
}

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
		if currentNode.get(letter) == nil {
			return nil
		}
		currentNode = currentNode.get(letter)
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
		if currentNode.get(letter) != nil {
			currentNode = currentNode.get(letter)
			continue
		}

		currentNode.set(letter)
		currentNode = currentNode.get(letter)
	}

	currentNode.set(terminateChar)
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
