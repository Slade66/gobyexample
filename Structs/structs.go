package main

import "fmt"

type person struct {
	name string
	age  int
}

func newPerson(name string) *person {
	return &person{name: name, age: 24}
}

func main() {

	fmt.Println(person{"bob", 20})
	fmt.Println(person{name: "alice", age: 30})
	fmt.Println(person{name: "fred"})
	fmt.Println(&person{name: "ana", age: 40})
	fmt.Println(newPerson("lyz"))
	
	lyz := person{"lyz", 24}
	fmt.Println(lyz.name)

	sp := &lyz
	fmt.Println(sp.name)

	cat := struct {
		name string
		color string
	} {"lili", "white"}
	fmt.Println(cat)
}
