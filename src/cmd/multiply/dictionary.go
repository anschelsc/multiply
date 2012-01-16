package main

const dataSize = 1024 * 1024

var dataSpace [dataSize]int

var here int // Index into dataSpace

type word struct {
	index     int // Positive means an index into dataSpace, negative into builtins
	immediate bool
}

var latest *word

var dictionary = make(map[string]*word)

// Define a new builtin word
func defineB(name string, immediate bool, b func()) {
	index := len(builtins)
	builtins = append(builtins, b)
	dictionary[name] = &word{-index, immediate}
	latest = dictionary[name]
}

// Define a new threaded word
// Should I append EXIT to the end? Is there any time I wouldn't want it?
func define(name string, immediate bool, cells []int) {
	copy(dataSpace[here:], cells)
	dictionary[name] = &word{here, immediate}
	here += len(cells)
	latest = dictionary[name]
}

// This is almost writing forth...
func colon(name string, immediate bool, words []string) {
	cells := make([]int, len(words))
	for i, w := range words {
		cells[i] = dictionary[w].index
	}
	define(name, immediate, cells)
}
