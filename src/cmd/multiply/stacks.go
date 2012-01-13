package main

type stack []int

func (s *stack) push(c int) {
	*s = append(*s, c)
}

func (s *stack) pop() int {
	ret := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return ret
}

var rstack, dstack *stack
