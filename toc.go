package toc

import (
	"strings"

	"github.com/yuin/goldmark/ast"
	"github.com/yuin/goldmark/parser"
	"github.com/yuin/goldmark/text"
	"github.com/yuin/goldmark/util"
)

var tocResultKey = parser.NewContextKey()

type tocTransformer struct{}

var defaultTransformer = new(tocTransformer)

func (t *tocTransformer) Transform(n *ast.Document, reader text.Reader, pc parser.Context) {
	var (
		inHeading   bool
		toc         = make([]Header, 0, 10)
		headingText strings.Builder
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
			if inHeading && entering {
				headingText.Write(n.Text(reader.Source()))
			}
		}
		return ast.WalkContinue, nil
	})
	pc.Set(tocResultKey, toc)
}
