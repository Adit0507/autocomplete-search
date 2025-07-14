package service

import (
	"github.com/Adit0507/autocomplete-search/internal/cache"
	"github.com/Adit0507/autocomplete-search/internal/models"
	"github.com/Adit0507/autocomplete-search/internal/trie"
)

type AutocompleteService struct {
	trie  *trie.Trie
	cache *cache.Cache
}

func NewAutoCompleteService(cacheSize int) *AutocompleteService {
	svc := &AutocompleteService{
		trie:  trie.NewTrie(),
		cache: cache.NewCache(cacheSize),
	}

	// sample data
	svc.trie.Insert("apple", 100)
	svc.trie.Insert("app", 80)
	svc.trie.Insert("apricot", 50)
	svc.trie.Insert("banana", 90)

	return svc
}

func (s *AutocompleteService) GetSuggestions(query string, limit int) []models.Suggestion {
	if suggestions, exists := s.cache.Get(query); exists {
		return suggestions
	}

	// query trie
	suggestions := s.trie.Search(query, limit)
	s.cache.Set(query, suggestions)	//storing in cache

	return suggestions
}
