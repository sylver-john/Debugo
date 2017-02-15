# Debugo
A Go var dump

##TODO

- [ ] Check if var is a struct
- [ ] Display Name and type of the var
- [ ] Display the length of the var as byte ? 
- [ ] Display all exported properties of the var

## Installation

Install Debugo using the "go get" command:
``go get github.com/sylver-john/Debugo/dump``

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

will print :
``
Var's name : a
Var's Type : A
Var's Kind : struct
Exported field :
field : Foo value : test type : string
field : Fuu value : 10 type : int
``

It must panic if you have a unexported filed in your struct and if you dump a interface.
