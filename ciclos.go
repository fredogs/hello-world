package main

import "fmt"

func main() {
/*	for i:=0; i<10; i++{
		fmt.Printf("Ronda: %d\n",i)
	}
	// Solo se requiere el ciclo for puesto que sus parametros son condiconales.
*/
	i:=0
	for {
		if i==2{
			i++
			continue
		}
		fmt.Println(i)
		i++
		if i>10{
			break
		}
	}
}