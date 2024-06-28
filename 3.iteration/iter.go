package iteration

import "fmt"

func Repeat(text string, times int) string {
	var concat string

	for i := 0; i < times; i++ {
		concat += text
	}

	return concat
}

func ExampleRepeat() {
	fmt.Println("ba", Repeat("na", 2))
	// Output: Hello
}
