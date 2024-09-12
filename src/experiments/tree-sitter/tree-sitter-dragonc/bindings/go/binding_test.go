package tree_sitter_dragonc_test

import (
	"testing"

	tree_sitter "github.com/tree-sitter/go-tree-sitter"
	tree_sitter_dragonc "github.com/tree-sitter/tree-sitter-dragonc/bindings/go"
)

func TestCanLoadGrammar(t *testing.T) {
	language := tree_sitter.NewLanguage(tree_sitter_dragonc.Language())
	if language == nil {
		t.Errorf("Error loading Dragonc grammar")
	}
}
