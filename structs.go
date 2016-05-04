package main

import "fmt"

type User struct {
		edad int
		nombre,apellido string
	}

func (this User) nombre_completo() string{
	return this.nombre + " " + this.apellido
}

func (this *User) set_name(n string){
	this.nombre = n
}

func main() {
	
	//var Frederick User

	//fmt.Println(User{nombre: "Frederick",apellido: "Gutierrez"})

	//Frederick := User{nombre: "Frederick",apellido: "Gutierrez"}
	//Erick := User{20,"",""}
	Fred := new(User)

	Fred.nombre = "Erick"
	Fred.set_name("Marcos")

	fmt.Println(Fred.nombre_completo())

}