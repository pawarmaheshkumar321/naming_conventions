package main

import "fmt"

type Person struct {
	firstname string
	lastname  string
	age       int
	address   Address
}

type Student struct {
	name   string
	rollno string
	mark1  int
	mark2  int
}

type Address struct {
	city, country string
}

func main() {

	p1 := Person{
		firstname: "Maheshkumar",
		lastname:  "Pawar",
		age:       22,
		address:   Address{city: "San Francisco", country: "USA"},
	}

	fmt.Println(p1.fullName())

	user1 := struct {
		username string
		password int
	}{
		username: "pawarmaheshkumar321@gmail.com",
		password: 123,
	}

	fmt.Printf("This is output %s %d\n", user1.username, user1.password)

	student1 := Student{
		name:   "Sumit",
		rollno: "10it",
		mark1:  10,
		mark2:  20,
	}

	fmt.Printf("Total marks of %s are %d\n", student1.name, student1.totalmarks())

	fmt.Printf("Incremented age %d\n", p1.incrementAgeByOne())
	fmt.Printf("Age %d\n", p1.age)

	fmt.Printf("City is %s and country %s\n", p1.address.city, p1.address.country)

}

func (p Person) fullName() string {
	return p.firstname + " " + p.lastname
}

func (s Student) totalmarks() int {
	var total int
	total = s.mark1 + s.mark2
	return total
}

func (p *Person) incrementAgeByOne() int {
	p.age++
	return p.age
}
