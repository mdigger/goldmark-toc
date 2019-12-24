package toc

import (
	"bytes"
	"regexp"
	"strings"
	"unicode"
	"unicode/utf8"

	"github.com/yuin/goldmark/ast"
	"github.com/yuin/goldmark/parser"
	"github.com/yuin/goldmark/text"
	"github.com/yuin/goldmark/util"
)

type tocTransformer struct{}

var (
	tocResultKey          = parser.NewContextKey()
	defaultTocTransformer = new(tocTransformer)
	reWords               = regexp.MustCompile(`[\S]+`)
)

var isWordsDivider = func(c rune) bool {
	return !unicode.IsLetter(c) && !unicode.IsNumber(c)
}

func (t *tocTransformer) Transform(n *ast.Document, reader text.Reader, pc parser.Context) {
	var (
		inHeading    bool
		toc          = make([]Header, 0, 100)
		headingText  strings.Builder
		words, chars int
	)
	ast.Walk(n, func(n ast.Node, entering bool) (ast.WalkStatus, error) {
		switch n.Kind() {
		case ast.KindHeading:
			inHeading = entering
			if !entering {
				var h = n.(*ast.Heading)
				var header = Header{
					Text:  headingText.String(),
					Level: h.Level,
				}
				headingText.Reset()
				if id, found := h.AttributeString("id"); found {
					header.ID = util.BytesToReadOnlyString(id.([]byte))
				}
				toc = append(toc, header)
			}
		case ast.KindText, ast.KindString:
			if !entering {
				break
			}
			var text = n.Text(reader.Source())
			if inHeading {
				headingText.Write(text)
			}
			chars += utf8.RuneCount(text)
			words += len(bytes.FieldsFunc(text, isWordsDivider))
			// fmt.Printf("%q\n", bytes.FieldsFunc(text, isPunct))
		}
		return ast.WalkContinue, nil
	})
	pc.Set(tocResultKey, &Info{
		Headers: toc,
		Words:   words,
		Chars:   chars,
	})
}
