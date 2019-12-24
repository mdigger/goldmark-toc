// Package toc adds the ability to get a table of contents to the
// goldmark parser.
package toc

import (
	"io"

	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/parser"
	"github.com/yuin/goldmark/text"
	"github.com/yuin/goldmark/util"
)

// ConverterFunc is markdown converting function.
type ConverterFunc = func(source []byte, writer io.Writer) (*Info, error)

// Markdown extends initialied goldmark.Markdown and return converter function.
func Markdown(m goldmark.Markdown) ConverterFunc {
	m.Parser().AddOptions(
		parser.WithAttribute(),
		parser.WithAutoHeadingID(),
		parser.WithASTTransformers(
			util.Prioritized(defaultTocTransformer, 1000),
		),
	)
	return func(source []byte, writer io.Writer) (*Info, error) {
		var ctx = parser.NewContext(parser.WithIDs(newIDs(Lang)))
		var doc = m.Parser().Parse(text.NewReader(source), parser.WithContext(ctx))
		if writer != nil {
			if err := m.Renderer().Render(writer, source, doc); err != nil {
				return nil, err
			}
		}
		if info, ok := ctx.Get(tocResultKey).(*Info); ok {
			return info, nil
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
func Convert(source []byte, writer io.Writer) (*Info, error) {
	return defaultMarkdown(source, writer)
}

// Statistic return maekdown document info.
func Statistic(source []byte) *Info {
	info, _ := defaultMarkdown(source, nil)
	return info
}
