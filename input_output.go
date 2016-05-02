package main

import(
	"fmt"
	"bufio"
	"os"
)

func main() {
/*
	edad := 22
	fmt.Printf("Mi edad es %d\n",edad)
	%d Entero
	%s Cadena
	%v Valor por defecto
	%t boolean
	%f Flotante (No tantos decimales)
	%e Cientifico
	%b Cientifico
*/
/*
	var nombre string
	fmt.Println("Coloca tu Nombre: ")
	fmt.Scanf("%s\n",&nombre)
	fmt.Printf("Hola %s\n",nombre)
*/
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Ingresa tu nombre: ")
	nombre,err := reader.ReadString('\n')
	if(err != nil){
		fmt.Println(err)
	}else{
		fmt.Println("Hola "+nombre)
	}
}