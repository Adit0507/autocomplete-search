package trie

import "github.com/Adit0507/autocomplete-search/internal/models"

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

func (t *Trie) Search(prefix string, limit int) []models.Suggestion {
	node := t.Root

	for _, char := range prefix {
		if _, exists := node.Children[char]; !exists {
			return nil
		}

		node = node.Children[char]
	}

	return t.collectSuggestions(node, prefix, limit)
}

func (t *Trie) collectSuggestions(node *Node, prefix string, limit int) []models.Suggestion {
	suggestions := []models.Suggestion{}
	if node.IsEnd {
		suggestions = append(suggestions, models.Suggestion{Text: prefix, Frequency: node.Frequency})
	}

	for char, child := range node.Children {
		childSuggestions := t.collectSuggestions(child, prefix+string(char), limit)
		suggestions = append(suggestions, childSuggestions...)

		if len(suggestions) >= limit {
			break
		}
	}

	return sortSuggestions(suggestions, limit)
}

func sortSuggestions(suggestions []models.Suggestion, limit int) []models.Suggestion {
	for i := 0; i < len(suggestions)-1; i++ {
		for j := i + 1; j < len(suggestions); j++ {
			if suggestions[i].Frequency < suggestions[j].Frequency {
				suggestions[i], suggestions[j] = suggestions[j], suggestions[i]
			}
		}
	}

	if len(suggestions) > limit {
		return suggestions[:limit]
	}

	return suggestions
}
