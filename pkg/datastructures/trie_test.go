package datastructures_test

import (
	"reflect"
	"testing"

	"github.com/fugu-chop/golang_dsa/pkg/datastructures"
)

func TestTrie_Insert(t *testing.T) {
	t.Parallel()

	t.Run("creates new child nodes from root", func(t *testing.T) {
		t.Parallel()

		trie := datastructures.Trie()
		trie.Insert("bus")

		letters := trie.Traverse(trie.Root(), []string{})
		expectedLetters := []string{"b", "u", "s", "*"}
		if !reflect.DeepEqual(letters, expectedLetters) {
			t.Fatalf("expected trie to contain: %v, got: %v", expectedLetters, letters)
		}
	})

	t.Run("appends to exising root", func(t *testing.T) {
		t.Parallel()

		trie := datastructures.Trie()
		trie.Insert("cat")
		trie.Insert("car")

		letters := trie.Traverse(trie.Root(), []string{})

		lettersMap := make(map[string]int)
		for _, letter := range letters {
			lettersMap[letter] += 1
		}

		expectedMap := map[string]int{
			"c": 1, "a": 1, "t": 1, "*": 2, "r": 1}
		if !reflect.DeepEqual(expectedMap, lettersMap) {
			t.Fatalf("expected trie to contain: %v, got: %v", expectedMap, lettersMap)
		}
	})

	t.Run("is no-op when all letters already exist", func(t *testing.T) {
		t.Parallel()

		trie := datastructures.Trie()
		trie.Insert("zoo")
		trie.Insert("zoo")

		letters := trie.Traverse(trie.Root(), []string{})
		expectedLetters := []string{"z", "o", "o", "*"}
		if !reflect.DeepEqual(letters, expectedLetters) {
			t.Fatalf("expected trie to contain: %v, got: %v", expectedLetters, letters)
		}
	})
}

func TestTrie_Search(t *testing.T) {
	t.Parallel()

	t.Run("searches when word exists", func(t *testing.T) {
		t.Parallel()

		trie := datastructures.Trie()
		trie.Insert("blimp")

		node := trie.Search("blimp")
		if node == nil {
			t.Fatal("expected node to exist, got nil")
		}
	})

	t.Run("doesn't return result when word partially exists", func(t *testing.T) {
		t.Parallel()

		trie := datastructures.Trie()
		trie.Insert("blim")

		node := trie.Search("blimp")
		if node != nil {
			t.Fatal("expected node to exist, got nil")
		}
	})

	t.Run("doesn't return result when word doesn't exist", func(t *testing.T) {
		t.Parallel()

		trie := datastructures.Trie()
		trie.Insert("animal")

		node := trie.Search("kangaroo")
		if node != nil {
			t.Fatal("expected node to exist, got nil")
		}
	})
}

func TestTrie_Autocomplete(t *testing.T) {
	t.Parallel()

	t.Run("handles partial prefixes", func(t *testing.T) {
		t.Parallel()

		trie := datastructures.Trie()
		trie.Insert("animal")
		trie.Insert("animus")
		trie.Insert("animosity")

		letters := trie.Autocomplete("ani")

		lettersMap := make(map[string]int)
		for _, entries := range letters {
			lettersMap[entries]++
		}

		expected := map[string]int{"animal": 1, "animus": 1, "animosity": 1}

		if !reflect.DeepEqual(expected, lettersMap) {
			t.Fatalf("expected autocomplete to contain: %v, got: %v", expected, lettersMap)
		}
	})

	t.Run("returns empty slice if prefix overflows", func(t *testing.T) {
		t.Parallel()

		trie := datastructures.Trie()
		trie.Insert("anima")

		prefix := "animal"
		letters := trie.Autocomplete(prefix)

		if len(letters) != 0 {
			t.Fatalf("expected no results, got: %v", letters)
		}
	})

	t.Run("returns empty slice if doesn't exist", func(t *testing.T) {
		t.Parallel()

		trie := datastructures.Trie()
		trie.Insert("animal")

		prefix := "bonk"
		letters := trie.Autocomplete(prefix)

		if len(letters) != 0 {
			t.Fatalf("expected no results, got: %v", letters)
		}
	})
}

func TestTrie_Traverse(t *testing.T) {
	t.Parallel()

	t.Run("traverses multiple structures", func(t *testing.T) {
		t.Parallel()

		trie := datastructures.Trie()
		trie.Insert("car")
		trie.Insert("cartouche")
		trie.Insert("cartographer")
		trie.Insert("cartwheels")

		letters := trie.Traverse(trie.Root(), []string{})

		lettersMap := make(map[string]int)
		for _, letter := range letters {
			lettersMap[letter] += 1
		}

		expected := map[string]int{
			"*": 4, "a": 2, "c": 2, "e": 4, "g": 1, "h": 3, "l": 1,
			"o": 1, "p": 1, "r": 3, "s": 1, "t": 1, "u": 1, "w": 1,
		}

		if !reflect.DeepEqual(expected, lettersMap) {
			t.Fatalf("expected traverse on trie to equal: %v, got: %v", expected, lettersMap)
		}
	})

	t.Run("when root node has no children, returns", func(t *testing.T) {
		t.Parallel()

		trie := datastructures.Trie()

		letters := trie.Traverse(trie.Root(), []string{})
		if len(letters) != 0 {
			t.Fatalf("expected traverse on empty trie to not have any children, got: %v", letters)
		}
	})

	t.Run("when node has no children, returns", func(t *testing.T) {
		t.Parallel()

		trie := datastructures.Trie()
		trie.Insert("car")

		node := trie.Search("car")
		if node == nil {
			t.Fatal("expected node to exist, got nil")
		}
		letters := trie.Traverse(node, []string{})
		expectedLetters := []string{"*"}
		if !reflect.DeepEqual(letters, expectedLetters) {
			t.Fatalf("expected trie to contain: %v, got: %v", expectedLetters, letters)
		}
	})
}
