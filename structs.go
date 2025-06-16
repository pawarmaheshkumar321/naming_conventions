package main

import "fmt"

type Person struct {
	firstname string
	lastname  string
	age       int
}

type Student struct {
	name   string
	rollno string
	mark1  int
	mark2  int
}

func main() {

	p1 := Person{
		firstname: "Maheshkumar",
		lastname:  "Pawar",
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

}

func (p Person) fullName() string {
	return p.firstname + " " + p.lastname
}

func (s Student) totalmarks() int {
	var total int
	total = s.mark1 + s.mark2
	return total
}
