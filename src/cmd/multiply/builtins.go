package main

func flag(b bool) int {
	if b {
		return -1
	}
	return 0
}

// Note that this copies the values TWICE:
// once in int->byte conversion, which is necessary,
// and once in []byte->string conversion which is unnecessary
// and can be avoided at some later date with unsafe tricks.
func uncounted(counted int) string {
	buf := make([]byte, dataSpace[counted])
	for i := range buf {
		buf[i] = byte(dataSpace[counted+1+i])
	}
	return string(buf)
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
	defineB("-", false, func() {
		x := dstack.pop()
		y := dstack.pop()
		dstack.push(y - x)
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
	// (CREATE) is the first cell of a CREATEd word;
	// the second is the address of the does> code if it exists.
	defineB("(CREATE)", false, func() {
		dstack.push(pc + 1)
		does := dataSpace[pc]
		if does == 0 {
			pc = rstack.pop() // EXIT
		} else {
			pc = does
		}
	})
	defineB("DEPTH", false, func() {
		dstack.push(len(*dstack))
	})
	defineB("DROP", false, func() {
		dstack.pop()
	})
	defineB("DUP", false, func() {
		dstack.push((*dstack)[len(*dstack)-1])
	})
	defineB("EXECUTE", false, func() {
		rstack.push(pc)
		pc = dstack.pop()
	})
	defineB("EXIT", false, func() {
		pc = rstack.pop()
	})
	defineB("FIND", false, func() {
		counted := dstack.pop()
		w, ok := dictionary[uncounted(counted)]
		if !ok {
			dstack.push(counted)
			dstack.push(0)
			return
		}
		dstack.push(w.index)
		if w.immediate {
			dstack.push(1)
		} else {
			dstack.push(-1)
		}
	})
	defineB("HERE", false, func() {
		dstack.push(here)
	})
	defineB("IMMEDIATE", false, func() {
		latest.immediate = true
	})
	defineB("INVERT", false, func() {
		dstack.push(^dstack.pop())
	})
	defineB("LATEST", false, func() {
		dstack.push(latest.index)
	})
	defineB("(LITERAL)", false, func() {
		dstack.push(dataSpace[pc])
		pc++
	})
	defineB("LSHIFT", false, func() {
		u := i2u(dstack.pop())
		x1 := dstack.pop()
		dstack.push(x1 << u)
	})
	defineB("NEGATE", false, func() {
		dstack.push(-dstack.pop())
	})
	defineB("OR", false, func() {
		dstack.push(dstack.pop() | dstack.pop())
	})
	defineB("R>", false, func() {
		dstack.push(dstack.pop())
	})
	defineB("R@", false, func() {
		dstack.push((*rstack)[len(*rstack)-1])
	})
	defineB("RSHIFT", false, func() {
		u := i2u(dstack.pop())
		x1 := dstack.pop()
		dstack.push(x1 >> u)
	})
	defineB("STATE", false, func() {
		dstack.push(len(dataSpace) - 1)
	})
	defineB("SWAP", false, func() {
		x := dstack.pop()
		y := dstack.pop()
		dstack.push(x)
		dstack.push(y)
	})
	defineB("XOR", false, func() {
		dstack.push(dstack.pop() ^ dstack.pop())
	})
}
