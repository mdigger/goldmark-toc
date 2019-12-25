package toc

import (
	"github.com/yuin/goldmark/ast"
	"github.com/yuin/goldmark/util"
)

// Header holds the data about a header.
type Header struct {
	Level int
	Text  string
	ID    string
}

// Headers return table of content from parsed markdown document.
func Headers(n ast.Node, source []byte) []Header {
	var toc = make([]Header, 0, 100)
	ast.Walk(n, func(n ast.Node, entering bool) (ast.WalkStatus, error) {
		if n.Kind() == ast.KindHeading && entering {
			var (
				h      = n.(*ast.Heading)
				text   = n.Text(source)
				header = Header{
					Text:  util.BytesToReadOnlyString(text),
					Level: h.Level,
				}
			)
			if id, found := h.AttributeString("id"); found {
				header.ID = util.BytesToReadOnlyString(id.([]byte))
			}
			toc = append(toc, header)
			return ast.WalkSkipChildren, nil
		}
		return ast.WalkContinue, nil
	})
	return toc
}

// // MarkdownHeaders writes table of content in markdown format.
// func MarkdownHeaders(w io.Writer, headers []Header, filename string) {
// 	var level, indent int
// 	for i, header := range headers {
// 		if i > 0 {
// 			if level < header.Level {
// 				indent++
// 			} else if indent > 0 {
// 				indent--
// 			}
// 		}
// 		level = header.Level
// 		prefix := strings.Repeat("  ", indent)
// 		var text string
// 		if header.ID != "" {
// 			text = fmt.Sprintf("[%s](%s#%s)", header.Text, filename, header.ID)
// 		} else {
// 			text = header.Text
// 		}
// 		fmt.Fprintf(w, "%s- %s\n", prefix, text)
// 	}
// }
