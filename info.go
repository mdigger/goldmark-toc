package toc

// Info describe table of contents, words & chars counters.
type Info struct {
	Headers      []Header
	Words, Chars int
}
