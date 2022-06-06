// 20/80
package main

import "fmt"

type person struct {
	name string
	age  int
}

func newPerson(name string) *person {
	p := person{name: name}
	p.age = 37
	return &p

}

func main() {
	fmt.Println(person{"Bob", 20})
	fmt.Println(person{name: "Alice", age: 11})
	fmt.Println(newPerson("Jason"))
	fmt.Println(&person{name: "hi", age: 30})
	s := person{"Sean", 102}
	fmt.Println(s.name)
	scp := &s
	fmt.Println(scp.age)

	scp.age = 103
	fmt.Println(scp.age)
}
