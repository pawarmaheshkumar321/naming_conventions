package main

import (
	"html/template"
	"os"
)

func main() {

	//tmpl := template.New("example")

	tmpl, err := template.New("example").Parse("Welcome, {{.name}}. How are you? ")

	if err != nil {
		panic(err)
	}

	data := map[string]interface{}{
		"name": "Maheshkumar",
	}

	err = tmpl.Execute(os.Stdout, data)

	tmpl1 := template.Must(template.New("example2").Parse("Welcome, {{.name}}. How are you? "))

	err = tmpl1.Execute(os.Stdout, data)

}
