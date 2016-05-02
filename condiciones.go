package main

import "fmt"

func main() {
	x := 10
	y := 50
/*
	== igual a
	!= Diferente de
	< menor que
	> mayor que
	>= mayor o igual que
	&& AND
	|| OR
*/
	if (x > y) {
		fmt.Printf("%d es mayor que %d\n",x,y)
	}else if x < y{
		fmt.Printf("%d es mayor que %d\n",y,x)
	}else{
		fmt.Printf("Son iguales")
	}
}