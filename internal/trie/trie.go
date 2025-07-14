package trie

type Node struct {
	Children  map[rune]*Node
	IsEnd     bool
	Frequency int
}

type Trie struct {
	Root *Node
}

func NewTrie() *Trie {
	return &Trie{Root: &Node{Children: make(map[rune]*Node)}}
}

func (t *Trie) Insert(word string, freq int) {
	node := t.Root
	for _, char := range word {
		if _, exists := node.Children[char]; !exists {
			node.Children[char] = &Node{Children: make(map[rune]*Node)}
		}

		node = node.Children[char]
	}

	node.IsEnd = true
	node.Frequency = freq
}
