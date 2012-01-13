package main

type stack []cell

func (s *stack) push(c cell) {
	*s = append(*s, c)
}

func (s *stack) pop() cell {
	ret := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return ret
}

var rstack, dstack *stack
