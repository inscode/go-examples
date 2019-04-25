package main

import (
	"go_dev/day04/example/demo-04"
	"fmt"
)

func main() {
	fmt.Println("test func new")
	demo_04.NewCode()

	fmt.Println("test func slice")
	demo_04.SliceCode()

	fmt.Println("test func recursive")
	demo_04.Recursive(0)

	fmt.Println("test func closure")
	demo_04.Closure()

	demo_04.TestIntSort()
	demo_04.TestStringSort()
	demo_04.SearchInt()
	demo_04.TestMap()
	demo_04.TestTwoMap()
	demo_04.TestMapping()
	demo_04.TestMapLock()

	arr := []int{11, 2, 23, 8, 31}
	demo_04.Bubble(arr)
	fmt.Println("--------------------------")

	s1 := make([]int, 10)
	s1 = []int{1, 2, 3, 4, 5, 6, 6, 7}
	fmt.Println(s1)

	s2 := s1[:len(s1)]
	fmt.Println(s2)

	demo_04.TestSlice()
}