// deret catalan dengan rumus
// (2n)! / (n + 1)!n!

package main

import (
	"fmt"
)

func main() {
	fmt.Println(catalan(7))  // 429
	fmt.Println(catalan(10)) // 16796
}

func catalan(n int) int {
	var faktorial func(int) int
	faktorial = func(n int) int {
		if n <= 1 {
			return 1
		}
		return n * faktorial(n-1)
	}
	return faktorial(2*n) / (faktorial(n+1) * faktorial(n))
}
