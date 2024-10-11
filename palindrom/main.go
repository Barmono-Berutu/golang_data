// cek apakah didalam list ada yang mengelompokkan sekumpulan kata yang termasuk ke dalam kata palindrom.

package main

import "fmt"

func main() {
	fmt.Println(groupPalindrome([]string{"katak", "civic", "kasur", "rusak"}))                  // [[katak, civic], kasur, rusak]
	fmt.Println(groupPalindrome([]string{"racecar", "seru", "kasur", "civic", "bilik", "kak"})) // [[racecar, civic, kak], seru, kasur, bilik]
	fmt.Println(groupPalindrome([]string{"masuk", "civic", "hahah", "garam"}))                  // [[civic, hahah], masuk, garam]
}

func groupPalindrome(words []string) []any {
	//TODO: your code here
	dataPalindrom := []string{}
	bukanPalindrom := []string{}

	for _, val := range words {
		isPalindrom := true
		for i := 0; i <= len(val)-1; i++ {
			if val[i] != val[len(val)-1-i] {
				isPalindrom = false
				break
			}
		}

		if isPalindrom {
			dataPalindrom = append(dataPalindrom, val)
		} else {
			bukanPalindrom = append(bukanPalindrom, val)
		}
	}
	hasil := []any{}
	hasil = append(hasil, dataPalindrom)
	for _, val := range bukanPalindrom {
		hasil = append(hasil, val)
	}
	return hasil
}
