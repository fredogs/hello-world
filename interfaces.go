package main

import "fmt"

type User interface{ 
	Permisos() int // 1 - 5
	Nombre() string
}

type Admin struct{
	name string
}
func (this Admin) Permisos() int{
	return 5
}
func (this Admin) Nombre() string{
	return this.name
}

type Editor struct{
	name string
}
func (this Editor) Permisos() int{
	return 3
}
func (this Editor) Nombre() string{
	return this.name
}

func auth(user User) string{
	permisos := user.Permisos()
	if permisos > 4 {
		return user.Nombre() + " tiene permisos de Administrador"
	}else if permisos ==3 {
		return user.Nombre() + " tiene permisos de Editor"
	}
	return ""
}

func main() {
//	admin := Admin{"Frederick"}
//	editor := Editor{"Erick"}

	usuarios := []User{Admin{"Frederick"},Editor{"Erick"}}
	for _,usuario := range usuarios{
		fmt.Println(auth(usuario))
	}


//	fmt.Println(auth(admin))
//	fmt.Println(auth(editor))
}

