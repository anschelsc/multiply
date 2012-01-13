package main

const dataSize = 1024 * 1024

var dataSpace [dataSize]cell

var here int // Index into dataSpace

type word struct {
	cell
	immediate bool
}

var dictionary = make(map[string]word)

// Define a new builtin word
func defineB(name string, immediate bool, b builtin) {
	dictionary[name] = word{b, immediate}
}

// Define a new threaded word
func define(name string, immediate bool, cells []cell) {
	copy(dataSpace[here:], cells)
	dictionary[name] = word{simple(here), immediate}
	here += len(cells)
}

// This is almost writing forth...
func colon(name string, immediate bool, words []string) {
	cells := make([]cell, len(words))
	for i, w := range words {
		cells[i] = dictionary[w]
	}
	define(name, immediate, cells)
}
