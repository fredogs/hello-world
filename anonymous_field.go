package main

import "fmt"

type Human struct{
	name string
}

func (this Human) hablar() string{
	return "Bla bla bla"
}

type Dummy struct{
	name string
}

type Tutor struct{
	Human
	Dummy
}

func (this Tutor) hablar() string{
	return this.Human.hablar() + " Jira"
}

func main() {
	tutor := Tutor{Human{"Fred"},Dummy{"Erick"}}

	fmt.Println(tutor.hablar())
}