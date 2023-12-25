package tmpl

import (
	"fmt"
	"strconv"
)

func ExampleNew() {
	template := New[int]("some.{var}", func(i int) map[string]string {
		return map[string]string{
			"{var}": strconv.Itoa(i),
		}
	})

	fmt.Println(template.Render(5))
	fmt.Println(template.Render(-1))

	// Output:
	// some.5
	// some.-1
}
