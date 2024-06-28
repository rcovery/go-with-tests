package integers

import (
	"fmt"
	"reflect"
)

func Add(x int, y int) int {
	var z int = x + y

	fmt.Println(reflect.TypeOf(z))

	return z
}
