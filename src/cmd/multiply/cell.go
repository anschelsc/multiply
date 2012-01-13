package main

type cell interface {
	value() int
	exec()
}

var pc cell

// simple is either numeric or a threaded word, i.e. an index into the data space
type simple int

func (s simple) value() int { return int(s) }
func (s simple) exec() {
	rstack.push(pc)
	pc = s
}

func flag(b bool) cell {
	if b {
		return ^simple(0)
	}
	return simple(0)
}

type builtin func()

// BUG: This means that two xt's with different actions may be "equal".
func (_ builtin) value() int { return 0 }
func (b builtin) exec()      { b() }

// CREATE should set does to -1
type created struct {
	data int
	does int
}

func (c created) value() int { return c.data }

func (c created) exec() {
	dstack.push(simple(c.data))
	if c.does != -1 {
		rstack.push(pc)
		pc = simple(c.does)
	}
}
