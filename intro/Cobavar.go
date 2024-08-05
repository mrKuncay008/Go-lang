package main

import "fmt"

/*
Variable bisa di deklarasikan di atas main
sebagai object
*/
var correct bool
var wrong bool
var object1 int
var object2 int

func main() {
	var a int = 1                   // Variable int dan nilainya
	var c string = "Hallo Ini Popy" // Variable String
	x := 2                          // Variable Infared (Tanpa tipeData)
	b := "Hallo Ini Bocilo"         // Variable Infared String
	var meng string                 // Mendeklarasikan Var string
	meng = "Pony"                   // Memberi value kepada var yang sudah di deklarasikan

	object1 = 2 // Mendklarasikan var int obj 1 dan obj 2
	object2 = 3
	correct = true
	wrong = false // boolean

	// print var boolean
	fmt.Println(correct, wrong)
	fmt.Println(meng)
	fmt.Println(a)
	fmt.Println(x)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println("Hasil nya : ", object1+object2) // menjumlah langsung di print obj 1 + obj 2
}
