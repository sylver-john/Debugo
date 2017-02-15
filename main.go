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
	fmt.Println("Var's name : " + name)
	fmt.Println("Var's Type :", reflect.TypeOf(a).Name())
	fmt.Println("Var's Kind :", reflect.ValueOf(a).Kind())
	fmt.Println("Exported field :")
    for i := 0; i < v.NumField(); i++ {
        field := v.Field(i)
		fmt.Println("field :", reflect.TypeOf(a).Field(i).Name, "value :", field.Interface(), "type :", field.Type())
    }
}