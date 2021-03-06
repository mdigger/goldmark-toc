# goldmark-toc

[![GoDoc](https://godoc.org/github.com/mdigger/goldmark-toc?status.svg)](https://godoc.org/github.com/mdigger/goldmark-toc)

[Goldmark](https://github.com/yuin/goldmark) extension for generating table of content.

```go
var source = []byte(`
# Title
paragraph text
## Section *1*
paragraph text
### Subsection *1.1*
paragraph text
## Section *2*
paragraph text
## Заголовок на русском
`)
info, err := toc.Convert(source, ioutil.Discard)
if err != nil {
	log.Fatal(err)
}
for _, header := range info.Headers {
	fmt.Printf("%+v\n", header)
}
```

```go
{Level:1 Text:Title ID:title}
{Level:2 Text:Section 1 ID:section-1}
{Level:3 Text:Subsection 1.1 ID:subsection-1-1}
{Level:2 Text:Section 2 ID:section-2}
{Level:2 Text:Заголовок на русском ID:zagolovok-na-russkom}
```
