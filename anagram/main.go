// cek kalimat anagram contohnya :
// Rusak => Kasur (anagram)
// rumah => harum (anagram)
// ubi => ibu (anagram)
// siap => sapi (anagram)
// mobil => lomba (bukan anagram)

package main

import (
	"fmt"
)

func main() {
	var kalimat1 string
	fmt.Print("Masukkan Kalimat 1 : ")
	fmt.Scan(&kalimat1)

	var kalimat2 string
	fmt.Print("Masukkan Kalimat 2 : ")
	fmt.Scan(&kalimat2)
	if anagram(kalimat1, kalimat2) {
		fmt.Println("Kalimat yang di input merupakan Anagram")
	} else {
		fmt.Println("Kalimat yang di input bukan Anagram")
	}

}

func anagram(kalimat1, kalimat2 string) bool {
	if len(kalimat1) != len(kalimat2) {
		return false
	}

	validasiKalimat := make(map[string]int)

	for i := 0; i < len(kalimat1); i++ {
		validasiKalimat[string(kalimat1[i])]++
		validasiKalimat[string(kalimat2[i])]--
	}

	for _, v := range validasiKalimat {
		if v != 0 {
			return false
		}
	}
	return true
}
