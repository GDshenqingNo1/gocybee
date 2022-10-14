package main

import "fmt"

func main() {

	var a, b string
	_, err := fmt.Scan(&a, &b)
	if err != nil {
		fmt.Println("Unknown name please try again")
	}
	var x, y int
	var a1 = 1
	var b1 = 1
	x = len(a)
	y = len(b)
	for i := 0; i < x; i++ {
		a1 = (int(a[i]) - 64) * a1
	}
	for j := 0; j < y; j++ {
		b1 = (int(b[j]) - 64) * b1
	}
	if a1%47 == b1%47 {
		fmt.Println("Go")
	} else {
		fmt.Println("Stay")
	}
}
