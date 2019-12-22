package toc_test

import (
	"fmt"
	"io/ioutil"
	"log"

	toc "github.com/mdigger/goldmark-toc"
)

func Example() {
	var source = []byte(`
# Title
paragraph text
## Section *1*
paragraph text
### Subsection *1.1*
paragraph text
## Section *2*
paragraph text
`)
	toc, err := toc.Convert(source, ioutil.Discard)
	if err != nil {
		log.Fatal(err)
	}
	for _, header := range toc {
		fmt.Printf("%+v\n", header)
	}
	// Output:
	// {ID:title Level:1 Text:Title}
	// {ID:section-1 Level:2 Text:Section 1}
	// {ID:subsection-1-1 Level:3 Text:Subsection 1.1}
	// {ID:section-2 Level:2 Text:Section 2}
}
