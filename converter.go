// Package toc adds the ability to get a table of contents to the
// goldmark parser.
package toc

import (
	"io"

	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/parser"
	"github.com/yuin/goldmark/text"
)

// ConverterFunc is markdown converting function.
type ConverterFunc = func(source []byte, writer io.Writer) ([]Header, error)

// Markdown extends initialied goldmark.Markdown and return converter function.
func Markdown(m goldmark.Markdown) ConverterFunc {
	m.Parser().AddOptions(
		parser.WithAttribute(),
		parser.WithAutoHeadingID(),
	)
	return func(source []byte, writer io.Writer) (toc []Header, err error) {
		doc := m.Parser().Parse(text.NewReader(source), WithIDs())
		toc = Headers(doc, source)
		if writer != nil {
			err = m.Renderer().Render(writer, source, doc)
		}
		return toc, err
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
