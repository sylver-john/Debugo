package dump

import(
	"fmt"
	"reflect"
)

func Dump(a interface{}, name string) {
	v := reflect.ValueOf(a)
	fmt.Println("Var's name : " + name)
	fmt.Println("Var's Type :", reflect.TypeOf(a).Name())
	fmt.Println("Var's Kind :", reflect.ValueOf(a).Kind())
	fmt.Println("Exported field :")
    for i := 0; i < v.NumField(); i++ {
        field := v.Field(i)
        fieldName := reflect.TypeOf(a).Field(i).Name
        fieldValue := field.Interface()
        fieldType := field.Type()
		fmt.Println("field :", fieldName, "value :", fieldValue, "type :", fieldType)
    }
}