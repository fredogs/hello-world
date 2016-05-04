package main

import "fmt"

func main() {
/*
	La función Copy copia el minimo de elementos en ambos arreglos
*/
	slice := []int{1,2,3,4,5}
	copia := make([]int,len(slice),cap(slice)*2)

	copy(copia,slice)

	fmt.Println(slice)
	fmt.Println(copia)

}