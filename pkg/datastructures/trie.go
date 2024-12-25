package datastructures

type Trie struct {
	Root *TrieNode
}

type TrieNode struct {
	Children map[rune]*TrieNode
	IsEnd    bool
}

func NewTrie() *Trie {
	return &Trie{
		Root: &TrieNode{
			Children: make(map[rune]*TrieNode),
			IsEnd:    false,
		},
	}
}

func (t *Trie) Insert(word string) {
	curr := t.Root

	for _, r := range word {
		if _, exists := curr.Children[r]; !exists {
			curr.Children[r] = &TrieNode{
				Children: make(map[rune]*TrieNode),
				IsEnd:    false,
			}
		}
		curr = curr.Children[r]
	}
	curr.IsEnd = true
}

func (t *Trie) Delete(word string) bool {
	return t.deleteHelper(t.Root, word, 0)
}

func (t *Trie) deleteHelper(node *TrieNode, word string, index int) bool {
	if index == len(word) {
		if !node.IsEnd {
			return false
		}
		node.IsEnd = false
		return len(node.Children) == 0
	}

	r := rune(word[index])
	child, exists := node.Children[r]
	if !exists {
		return false
	}

	if shouldDeleteChild := t.deleteHelper(child, word, index+1); shouldDeleteChild {
		delete(node.Children, r)
		return len(node.Children) == 0 && !node.IsEnd
	}

	return false
}

func (t *Trie) Search(word string) bool {
	node := t.searchPrefix(word)
	return node != nil && node.IsEnd
}

func (t *Trie) StartsWith(prefix string) bool {
	return t.searchPrefix(prefix) != nil
}

func (t *Trie) searchPrefix(prefix string) *TrieNode {
	curr := t.Root

	for _, r := range prefix {
		if _, exists := curr.Children[r]; !exists {
			return nil
		}
		curr = curr.Children[r]
	}
	return curr
}
