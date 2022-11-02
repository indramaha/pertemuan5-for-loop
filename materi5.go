//Contoh 1. Cetak Barisan
/*
buatlah algoritma untuk mencetak barisan bilangan dari n ke m. Dengan n < m, dan n, m merupakan bilangan bulat positif. Misalkan, jika masukan n = 4, m = 8, maka bilangan yang dicetak 4 5 6 7 8

masukan terdiri dari 2 angka. Angka pertama menyatakan n dan angka kedua menyatakan m.

keluaran berupa barisan bilangan dari n hingga m.

contoh masukan dan keluaran
_______________________________________________________________________________
No	|			Masukan					|			Keluaran
1	|		1	10						|	1 2 3 4 5 6 7 8 9 10
2	|		4	8						|	4 5 6 7 8
3	|		9	14						|	9 10 11 12 13 14
_______________________________________________________________________________

*/

//Pseudocode
/*
program Cetak_Barisan
kamus
	n, m, i : integer
algoritma
	input(n, m)
	for i <- n to m do
		output(i)
	endfor
endprogram
*/

package main

import "fmt"

func main() {
	var n, m, i int
	fmt.Scan(&n, &m)
	for i = n; i <= m; i++ {
		fmt.Printf("%d", i)
	}
}
