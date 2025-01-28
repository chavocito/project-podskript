package main

import (
	"os"
	"text/template"
)

type Pets struct {
	Name   string
	Age    int
	Sex    string
	Breed  string
	Intact bool
}

func main() {
	dogs := []Pets{
		{"Bella", 3, "F", "Poodle", false},
		{"Max", 2, "M", "Golden Retriever", true},
		{"Lucy", 4, "F", "Labrador", false},
	}

	var tmplFile = "./views/pets.tmpl"
	tmpl, err := template.ParseFiles(tmplFile)
	if err != nil {
		panic(err)
	}

	err = tmpl.Execute(os.Stdout, dogs)
	if err != nil {
		panic(err)
	}
}
