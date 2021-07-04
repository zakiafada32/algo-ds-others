package main

import "fmt"

// AlphabetSize is the number of possible characters in the trie
const AlphabetSize = 26

// Node represents each node in the trie
type Node struct {
	children [AlphabetSize]*Node
	isEnd    bool
}

// Trie represent a true and has a pointer to the root node
type Trie struct {
	root *Node
}

// InitTrie will create a new trie
func InitTrie() *Trie {
	return &Trie{
		root: &Node{},
	}
}

// Insert will take in a word and add it to the trie
func (t *Trie) Insert(w string) {
	currentNode := t.root
	for _, v := range w {
		charIndex := v - 'a'
		if currentNode.children[charIndex] == nil {
			currentNode.children[charIndex] = &Node{}
		}
		currentNode = currentNode.children[charIndex]
	}
	currentNode.isEnd = true
}

// Search will take in a word and return true is that word is included in the trie
func (t *Trie) Search(w string) bool {
	currentNode := t.root
	for _, v := range w {
		charIndex := v - 'a'
		if currentNode.children[charIndex] == nil {
			return false
		}
		currentNode = currentNode.children[charIndex]
	}

	return currentNode.isEnd
}

func main() {
	myTrie := InitTrie()

	toAdd := []string{
		"aragorn",
		"aragon",
		"argon",
		"eragon",
		"oregon",
		"oregano",
		"oreo",
	}

	for _, v := range toAdd {
		myTrie.Insert(v)
	}

	fmt.Println(myTrie.Search("oreo"))
	fmt.Println(myTrie.Search("or"))
	fmt.Println(myTrie.Search("some"))
}
