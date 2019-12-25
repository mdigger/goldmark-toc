package toc

import (
	"fmt"

	"github.com/gosimple/slug"
	"github.com/yuin/goldmark/ast"
	"github.com/yuin/goldmark/parser"
	"github.com/yuin/goldmark/util"
)

// Lang define the default language for generating slug id fo headers.
var Lang = ""

type ids struct {
	lang   string
	values map[string]struct{}
}

// NewIDs return initialized ID generator for goldmark parser context.
func NewIDs(lang string) parser.IDs {
	return &ids{
		lang:   lang,
		values: make(map[string]struct{}),
	}
}

func (s *ids) Generate(value []byte, kind ast.NodeKind) []byte {
	var (
		slugStr = slug.MakeLang(
			util.BytesToReadOnlyString(value), s.lang)
		counter int
	)
	if slugStr == "" {
		slugStr = "id"
	}
	var result = slugStr
	for {
		if _, ok := s.values[result]; !ok {
			s.values[result] = struct{}{}
			return util.StringToReadOnlyBytes(result)
		}
		counter++
		result = fmt.Sprintf("%s-%d", slugStr, counter)
	}
}

func (s *ids) Put(value []byte) {
	s.values[util.BytesToReadOnlyString(value)] = struct{}{}
}

// WithIDs return new initializer parser option with ID generator.
func WithIDs() parser.ParseOption {
	return parser.WithContext(
		parser.NewContext(
			parser.WithIDs(NewIDs(Lang)),
		),
	)
}
