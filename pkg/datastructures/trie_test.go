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
		expectedLetters := []string{"c", "a", "t", "*", "r", "*"}
		if !reflect.DeepEqual(letters, expectedLetters) {
			t.Fatalf("expected trie to contain: %v, got: %v", expectedLetters, letters)
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

		if node.Get("*") == nil {
			t.Fatal("expected final character to be terminated by *")
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

		expected := []string{"animal", "animus", "animosity"}
		letters := trie.Autocomplete("ani")

		if !reflect.DeepEqual(expected, letters) {
			t.Fatalf("expected autocomplete to contain: %v, got: %v", expected, letters)
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
