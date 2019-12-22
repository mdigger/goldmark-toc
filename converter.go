// Package withtoc adds the ability to get a table of contents to the
// goldmark parser.
package withtoc

import (
	"io"

	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/parser"
	"github.com/yuin/goldmark/util"
)

// ConverterFunc is markdown converting function.
type ConverterFunc = func(source []byte, writer io.Writer) ([]Header, error)

// Markdown extends initialied goldmark.Markdown and return converter function.
func Markdown(m goldmark.Markdown) ConverterFunc {
	m.Parser().AddOptions(
		parser.WithAttribute(),
		parser.WithAutoHeadingID(),
		parser.WithASTTransformers(
			util.Prioritized(defaultTransformer, 1000),
		),
	)
	return func(source []byte, writer io.Writer) ([]Header, error) {
		var ctx = parser.NewContext(parser.WithIDs(newIDs(Lang)))
		if err := m.Convert(source, writer, parser.WithContext(ctx)); err != nil {
			return nil, err
		}
		if toc, ok := ctx.Get(tocResultKey).([]Header); ok {
			return toc, nil
		}
		return nil, nil
	}
}

// New return markdown converter with table of content support.
func New(options ...goldmark.Option) ConverterFunc {
	return Markdown(goldmark.New(options...))
}

var defaultMarkdown = Markdown(goldmark.New())

// Convert from markdown to html with default options and return TOC.
func Convert(source []byte, writer io.Writer) ([]Header, error) {
	return defaultMarkdown(source, writer)
}
