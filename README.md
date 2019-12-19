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
if err != nil {
	log.Fatal(err)
}
for _, header := range toc {
	fmt.Printf("%+v\n", header)
}
```

```go
{ID:toc:01 Level:1 Text:Title}
{ID:toc:02 Level:2 Text:Section 1}
{ID:toc:03 Level:3 Text:Subsection 1.1}
{ID:toc:04 Level:2 Text:Section 2}
```