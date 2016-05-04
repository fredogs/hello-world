package main

import "fmt"

func main() {
//	var matriz []int
//	matriz := []int {1,2,3,4}
/*	matriz := []int{1,2}
	if matriz == nil {
		fmt.Println("Esta vacio")
	}else{
		fmt.Println(len(matriz))
	}
*/
	arreglo := [3]int{1,2,3}
	slice := arreglo[0:2]
	fmt.Println(slice)
}