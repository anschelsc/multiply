package main

func flag(b bool) int {
	if b {
		return -1
	}
	return 0
}

var builtins = []func(){nil} // Need dummy element at 0

// Important builtins, in mostly ASCIIbetical order:

func initB() {
	defineB("!", false, func() {
		address := dstack.pop()
		val := dstack.pop()
		dataSpace[address] = val
	})
	defineB("*", false, func() {
		x := dstack.pop()
		y := dstack.pop()
		dstack.push(x * y)
	})
	defineB("+", false, func() {
		x := dstack.pop()
		y := dstack.pop()
		dstack.push(x + y)
	})
	defineB("/MOD", false, func() {
		x := dstack.pop()
		y := dstack.pop()
		dstack.push(y / x)
		dstack.push(y % x)
	})
	defineB("<", false, func() {
		x := dstack.pop()
		y := dstack.pop()
		dstack.push(flag(y < x))
	})
	defineB("=", false, func() {
		x := dstack.pop()
		y := dstack.pop()
		dstack.push(flag(y == x))
	})
	defineB(">", false, func() {
		x := dstack.pop()
		y := dstack.pop()
		dstack.push(flag(y > x))
	})
	defineB(">R", false, func() {
		rstack.push(dstack.pop())
	})
	defineB("@", false, func() {
		dstack.push(dataSpace[dstack.pop()])
	})
	defineB("ALLOT", false, func() {
		here += dstack.pop()
	})
	defineB("AND", false, func() {
		x := dstack.pop()
		y := dstack.pop()
		dstack.push(x & y)
	})
	defineB("BRANCH", false, func() {
		pc = dataSpace[pc]
	})
	defineB("0BRANCH", false, func() {
		if dstack.pop() == 0 {
			pc = dataSpace[pc]
		} else {
			pc++ // Skip the branch-address
		}
	})
}
