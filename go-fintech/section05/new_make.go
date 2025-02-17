package main

import "fmt"

func main() {
	// new
	var p *int = new(int)
	fmt.Println(p)  // 0x1400009c008
	fmt.Println(*p) // 0
	*p++
	fmt.Println(*p) // 1

	var p2 *int
	fmt.Println(p2) // <nil>
	// *p2++ -> panic
	// fmt.Println(p2)

	// make
	fmt.Println(" ========= make ==========")
	s := make([]int, 0)
	fmt.Printf("%T\n", s) // []int

	m := make(map[string]int)
	fmt.Printf("%T\n", m) // map[string]int

	ch := make(chan int)
	fmt.Printf("%T\n", ch) // chan int

	var p3 *int = new(int)
	fmt.Println(p3)

	var a *int
	fmt.Println(a) // nil

	var b *[]int
	fmt.Println(b) // nil

	var c *map[string]int
	fmt.Println(c)
	//fmt.Println(*c) // nil

	var st *struct{}
	fmt.Println(st) // nil
}
