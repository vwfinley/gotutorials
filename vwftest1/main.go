package main

import (
	"fmt"
)

func main() {
	y := 5
	var result = add(y, 5)

	fmt.Printf("hello world! x = %d\n", result)
}

func add(a, b int) int {
	return a + b
}

func sub() {

	var a []int           // declared, variable size
	a = []int{1, 2, 3, 4} // implicit size assignment

	var b [4]int           // declared fixed size
	b = [4]int{1, 2, 3, 4} // explicit size assignment

	c := []int{1, 2, 3, 4}    // implicit declaration, implicit size
	d := [...]int{1, 2, 3, 4} // implicit declaration, implicit size
	e := [4]int{1, 2, 3, 4}   // implicit declaration, explicit size

	var f = []int{1, 2, 3, 4}    // implicit declaration, implicit size
	var g = [...]int{1, 2, 3, 4} // implicit declaration, implicit size
	var h = [4]int{1, 2, 3, 4}   // implicit declaration, explicit size

	var i []int = []int{1, 2, 3, 4} // explicit declaration, implicit size
	//	var j []int = [...]int{1, 2, 3, 4} // [...] is invalid here
	var k [4]int = [4]int{1, 2, 3, 4} // explicit declaration, explicit size

	fmt.Print(a)
	fmt.Print(b)

	fmt.Print(c)
	fmt.Print(d)
	fmt.Print(e)
	fmt.Print(f)
	fmt.Print(g)
	fmt.Print(h)
	fmt.Print(i)
	//	fmt.Print(j)
	fmt.Print(k)
}
