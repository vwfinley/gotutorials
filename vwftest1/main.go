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

	var a []int  // declared, variable size
	var b [4]int // declared fixed size

	a = []int{1, 2, 3, 4}
	b = [4]int{1, 2, 3, 4}

	j := []int{1, 2, 3, 4}    // implicit declaration, implicit size
	k := [...]int{1, 2, 3, 4} // implicit declaration, implicit size
	l := [4]int{1, 2, 3, 4}   // implicit declaration, explicit size

	var m = []int{1, 2, 3, 4}    // implicit declaration, implicit size
	var n = [...]int{1, 2, 3, 4} // implicit declaration, implicit size
	var o = [4]int{1, 2, 3, 4}   // implicit declaration, explicit size

	var p []int = []int{1, 2, 3, 4} // explicit declaration, implicit size
	//	var q []int = [...]int{1, 2, 3, 4} // [...] is invalid here
	var r [4]int = [4]int{1, 2, 3, 4} // explicit declaration, explicit size

	fmt.Print(a)
	fmt.Print(b)

	fmt.Print(j)
	fmt.Print(k)
	fmt.Print(l)
	fmt.Print(m)
	fmt.Print(n)
	fmt.Print(o)
	fmt.Print(p)
	//	fmt.Print(q)
	fmt.Print(r)
}
