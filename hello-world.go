package main

import (
	"fmt"
	"reflect"
)

/*
  var (
	name, course string  // name of subscriber
	module       float64 // step in course
)
*/

var name, course, module = "Mike", "Learn Go", 2.6

func main() {
	fmt.Println("Name is set to", name, "and is of type", reflect.TypeOf(name))
	fmt.Println("Module is set to", module, "and is of type", reflect.TypeOf(module))

	a := 10.00000
	b := 3

	c := int(a) + b

	fmt.Println("C has value:", c, "and is of type", reflect.TypeOf(c))

	pointerName := "Mike"
	ptr := &pointerName

	fmt.Println("Name is", pointerName, "and is of type", reflect.TypeOf(pointerName))
	fmt.Println("Memory address of name is", ptr)
	fmt.Println("Value of pointer variable is", *ptr)
}
