# Debugo
A Go var dump

##TODO

- [ ] Check if var is a struct
- [ ] Display Name and type of the var
- [ ] Display the length of the var as byte ? 
- [ ] Display all exported properties of the vae

## Installation

Install Debugo using the "go get" command:
``go get go get github.com/sylver-john/Debugo/dump``

## Usage

Import the package with :
```go 
import(
	"github.com/sylver-john/Debugo/dump"
)
```

dump a var like :
```go 
type A struct {
   Foo string
   Fuu int
}

func main() {
	a := A{"test", 10}
	dump.Dump(a, "a")
}
```

It must panic if you have a unexported filed in your struct and if you dump a interface.
