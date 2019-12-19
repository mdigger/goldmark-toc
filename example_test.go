package converter_test

import (
	"fmt"
	"io/ioutil"
	"log"

	converter "github.com/mdigger/goldmark-toc"
)

func Example() {
	var markdown = converter.New() // initialize converter
	var source = []byte(`
# Title
paragraph text
## Section 1
paragraph text
## Subsection 1.1
paragraph text
## Section 2
paragraph text
`)
	toc, err := markdown(source, ioutil.Discard)
	if err != nil {
		log.Fatal(err)
	}
	for _, header := range toc {
		fmt.Println(header)
	}
	// Output:
	// {toc:01 1 Title}
	// {toc:02 2 Section 1}
	// {toc:03 2 Subsection 1.1}
	// {toc:04 2 Section 2}
}
