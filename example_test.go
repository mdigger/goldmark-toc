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
## Section *1*
paragraph text
### Subsection *1.1*
paragraph text
## Section *2*
paragraph text
`)
	toc, err := markdown(source, ioutil.Discard)
	if err != nil {
		log.Fatal(err)
	}
	for _, header := range toc {
		fmt.Printf("%+v\n", header)
	}
	// Output:
	// {ID:toc:01 Level:1 Text:Title}
	// {ID:toc:02 Level:2 Text:Section 1}
	// {ID:toc:03 Level:3 Text:Subsection 1.1}
	// {ID:toc:04 Level:2 Text:Section 2}
}
