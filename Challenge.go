package main

import(
	"fmt"
	//"math/rand"
)


func main() {
/*	x := rand.Intn(15)

	slice := make([]int,0,0)
	for i:=0; i<x; i++ {
		slice = append(slice,rand.Intn(10))
	}
	fmt.Println(slice)
*/
	channel := make(chan int)
	x := [10]int{1,2,3,4,5,6,7,8,9,10}
	//n := len(x)
	var sum int
	var c int
	go func(channel chan int){
		for i:=0;i<10;i++{
			sum = sum + x[i]
			channel <- i
				
		}
		c = <-channel
	}(channel)

	fmt.Println(c,sum)

}

	
	
