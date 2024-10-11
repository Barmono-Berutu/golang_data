package main

import "fmt"

func main() {
	n := 5 // hasil faktorial 120
	hasil := 1
	for i := n; i >= 1; i-- {
		hasil *= i
	}
	fmt.Println(hasil)
}
