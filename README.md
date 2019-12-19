# goldmark-toc

[![GoDoc](https://godoc.org/github.com/mdigger/goldmark-toc?status.svg)](https://godoc.org/github.com/mdigger/goldmark-toc)

[Goldmark](https://github.com/yuin/goldmark) extension for generating table of content.

```go
var markdown = converter.New() // initialize converter
var source = []byte(`
# Title
paragraph text
## Section 1
paragraph text
### Subsection 1.1
paragraph text
## Section 2
paragraph text
`)
toc, err := markdown(source, os.Stdout)
```