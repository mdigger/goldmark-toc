package toc

import (
	"fmt"
	"io"
	"strings"
)

// Header holds the data about a header.
type Header struct {
	Level int
	Text  string
	ID    string
}

// MarkdownHeaders writes table of content in markdown format.
func MarkdownHeaders(w io.Writer, headers []Header, filename string) {
	var level, indent int
	for i, header := range headers {
		if i > 0 {
			if level < header.Level {
				indent++
			} else if indent > 0 {
				indent--
			}
		}
		level = header.Level
		prefix := strings.Repeat("  ", indent)
		var text string
		if header.ID != "" {
			text = fmt.Sprintf("[%s](%s#%s)", header.Text, filename, header.ID)
		} else {
			text = header.Text
		}
		fmt.Fprintf(w, "%s- %s\n", prefix, text)
	}
}
