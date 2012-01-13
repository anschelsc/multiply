package main

var pc int

func exec(index int) {
	if index >= 0 {
		rstack.push(pc)
		pc = dataSpace[index]
	} else {
		builtins[-index]()
	}
}

func run() {
	for {
		pc++         // When exec is run,
		exec(pc - 1) // pc is the next instruction.
	}
}
