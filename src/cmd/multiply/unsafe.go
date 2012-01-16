package main

import "unsafe"

func i2u(i int) uint {
	return *(*uint)(unsafe.Pointer(&i))
}