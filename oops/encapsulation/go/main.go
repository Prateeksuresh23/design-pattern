package main

import (
	"fmt"

	"github.com/practice/design-pattern/oops/encapsulation/go/data"
)

func main() {
	myObject := data.NewUnexportedFields("value_for_field_1", "value_for_field_2")
	fmt.Println(myObject)
}
