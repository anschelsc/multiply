package main

// Important builtins, in mostly ASCIIbetical order:

func initB() {
	defineB("!", false, func() {
		address := dstack.pop().value()
		val := dstack.pop()
		dataSpace[address] = val
	})
	defineB("*", false, func() {
		x := dstack.pop().value()
		y := dstack.pop().value()
		dstack.push(simple(x * y))
	})
	defineB("+", false, func() {
		x := dstack.pop().value()
		y := dstack.pop().value()
		dstack.push(simple(x + y))
	})
	defineB("/MOD", false, func() {
		x := dstack.pop().value()
		y := dstack.pop().value()
		dstack.push(simple(y / x))
		dstack.push(simple(y % x))
	})
	defineB("<", false, func() {
		x := dstack.pop().value()
		y := dstack.pop().value()
		dstack.push(flag(y < x))
	})
	defineB("=", false, func() {
		x := dstack.pop().value()
		y := dstack.pop().value()
		dstack.push(flag(y == x))
	})
	defineB(">", false, func() {
		x := dstack.pop().value()
		y := dstack.pop().value()
		dstack.push(flag(y > x))
	})
	defineB(">R", false, func() {
		rstack.push(dstack.pop())
	})
	defineB("@", false, func() {
		dstack.push(dataSpace[dstack.pop().value()])
	})
	defineB("ALLOT", false, func() {
		here += dstack.pop().value()
	})
	defineB("AND", false, func() {
		x := dstack.pop().value()
		y := dstack.pop().value()
		dstack.push(simple(x & y))
	})
	defineB("BRANCH", false, func() {
		pc = dataSpace[pc.value()]
	})
	defineB("0BRANCH", false, func() {
		if dstack.pop().value() == 0 {
			pc = simple(dataSpace[pc.value()])
		} else {
			pc = simple(pc.value()+1)
		}
	})
}
