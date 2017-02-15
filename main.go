package main

import(
	"./dump"
)
type A struct {
   Foo string
   Fuu int
}

func main() {
	a := A{"test", 10}
	dump.Dump(a, "a")
}