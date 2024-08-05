package main

import "fmt"

/*
variable const yang nilainya tidak bisa
di rubah
*/
const (
	first = "Popy"
	last  = "ndeyou"
)

func main() {
	fmt.Println("Nama nya: ", first, last)
}
