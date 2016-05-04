package main

import "fmt"

func main() {
	slice := make([]int,0,2)
//	Puntero, Longitud, Capacidad
	slice = append(slice,2)
	fmt.Println(slice)
}