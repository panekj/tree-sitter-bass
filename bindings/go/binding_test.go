package tree_sitter_bass_test

import (
	"testing"

	tree_sitter "github.com/smacker/go-tree-sitter"
	"github.com/tree-sitter/tree-sitter-bass"
)

func TestCanLoadGrammar(t *testing.T) {
	language := tree_sitter.NewLanguage(tree_sitter_bass.Language())
	if language == nil {
		t.Errorf("Error loading Bass grammar")
	}
}
