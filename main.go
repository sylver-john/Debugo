package main

import(
	"fmt"
	"reflect"
)

type A struct {
   Foo string
   Fuu int
}

func main() {
	a := A{"test", 10}
	dump(a, "a")
}

func dump(a interface{}, name string) {
	v := reflect.ValueOf(a)
	fmt.Print("Var's name : " + name + "\n")
	fmt.Print("Var's Type : ", reflect.TypeOf(a).Name(), "\n")
	fmt.Print("Var's Kind : ", reflect.ValueOf(a).Kind(), "\n")
    values := make([]interface{}, v.NumField())
    for i := 0; i < v.NumField(); i++ {
        values[i] = v.Field(i)
		fmt.Print(values[i], "\n")
		fmt.Print(reflect.TypeOf(values[i]).Name(), "\n")
    }
}