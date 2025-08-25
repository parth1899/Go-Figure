package main

import "fmt"

func main() {

	// Array basics
	// var a [4]int
	// a[0] = 1
	// i := a[0]
	// fmt.Println("i: ", i)
	// fmt.Println("a: ", a)
	// // i:  1
	// // a:  [1 0 0 0]

	// // Array Literals
	// b := [2]string{"Parth", "Kalani"}
	// fmt.Println(b)

	// c := [...]string{"Parth", "Kalani"} // compiler counts the number of elements
	// fmt.Println(c)

	// Slices
	
	// 1. Slice Literal
	numbers := []int{1, 2, 3, 4, 5}
	fmt.Println(numbers)

	// 2. Slice with make
	s := make([]byte, 5)
	fmt.Println("s: ", s)
	fmt.Println("len: ", len(s))
	fmt.Println("cap: ", cap(s))

	// 3. Slicing an array
	b := []byte{'p', 'a', 'r', 't', 'h'};

	fmt.Println(b[1:4]) // [97 114 116]
	fmt.Println(b[:2])
	fmt.Println(b[2:])
	fmt.Println(b[:])
	// In Go, a byte is just an alias for uint8.
	// When you fmt.Println a []byte, it prints the decimal values of each byte.

	// 4. Array to Slice
	x := [3]string{"Bhramha", "Vishnu", "Mahesh"}
	s2 := x[:]
	fmt.Println(s2)

	// 5. Modifying Shared Backing Array
	d := []byte{'r', 'o', 'a', 'd'}
	e := d[2:]
	fmt.Println("e: ", string(e))

	e[1] = 'm';
	fmt.Println("e: ", string(e))
	fmt.Println("d: ", string(d))

	// 6. Growing Slices - Manual Reallocation & Copy Function
	s = []byte{1, 2, 3}
	// fmt.Println(cap(s))
	t := make([]byte, len(s), (cap(s) + 1)*2)
	copy(t, s) // dest, src
	s = t
	fmt.Println("s: ", s)
	// fmt.Println(cap(s))

	// 7. Custom AppendByte Function
	p := []byte{2, 3, 5}
	p = AppendByte(p, 7, 10, 12)
	fmt.Println(p)
	
	// 8. Built In Append function
	a := make([]int, 1)
	fmt.Println(a)

	a = append(a, 1, 2, 3)
	fmt.Println(a)

	b1 := []string{"Complexity", "is the", "enemy"}
	c1 := []string{"of", "security"}
	b1 = append(b1, c1...)
	fmt.Println(b1)

	// Example- Appending with a Condition (Filter)
	nums := []int{1, 2, 3, 4, 5, 6}
	nums = Filter(nums, func(n int) bool{ return n%2 == 0 })
	fmt.Println(nums)
	
}


func AppendByte(slice []byte, data ...byte) []byte {
	m := len(slice)
	n := m + len(data)

	if n > cap(slice) {
		// reallocate
		newSlice := make([]byte, (n+1)*2)
		copy(newSlice, slice)
		slice = newSlice
	}

	slice = slice[:n]
	copy(slice[m:n], data)
	return slice
}

func Filter(s []int, fn func(int) bool) []int {
	var p []int
	for _, v := range s {
		if(fn(v)) {
			p = append(p, v)
		}
	}
	return p
}