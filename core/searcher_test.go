package core

import (
	"testing"
)

func TestNewSearcherWithNoId(t *testing.T) {
	searcher := NewSearcher(SearcherOptions{})
	if searcher.Id != "searcher" {
		t.Errorf("Expected searcher.Id to be 'searcher', got %s", searcher.Id)
	}
}

func TestNewSearcherWithGivenId(t *testing.T) {
	searcher := NewSearcher(SearcherOptions{Id: "mysearcher"})
	if searcher.Id != "mysearcher" {
		t.Errorf("Expected searcher.Id to be 'mysearcher', got %s", searcher.Id)
	}
}
