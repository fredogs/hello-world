package main

import(
	"fmt"
	"time"
	"strings"
)

func main() {
	go mi_nombre_lento("Fred")
	fmt.Println("Queee Aburridooo")
	var wait string
	fmt.Scanln(&wait)
}

func mi_nombre_lento(name string){
	letras := strings.Split(name,"")

	for _,letra := range(letras){
		time.Sleep(1000 * time.Millisecond)
		fmt.Println(letra)
	}

}