package main

import "fmt"

type Curso struct {
	id int
	nombre string
	profesor string
}

func main()  {
	//Declaring an Array
	/*
	var cursos[3] string

	// Fill the array Classic Way
	cursos[0] = "Go"
	cursos[1] ="JavaScript"
	cursos[2] = "C sharp"
	for _, value := range cursos{
		fmt.Println(value)
	}
	*/
	cursos := [4] string {"Go", "Python", "Javascript", "C#"}

	// Using for Range
	for _, value := range cursos{
		fmt.Println(value)
	}

	// Using classic FOR
	for i :=0; i < len(cursos); i++ {
		fmt.Println(cursos[i])

	}

	// Declaring  a slice
	var cursos2[] string

	cursos2 = append(cursos2, "Python")
	cursos2 = append(cursos2,"GO")
	cursos2 = append(cursos2, "C #")
	cursos2 = append(cursos2,"JavaScript")
	fmt.Println(cursos2)

	for _, value := range(cursos2) {
		fmt.Println(value)
	}

	// Short way of create an slice
	cursos3 := make([]string, 0)
	cursos3 = []string{"Go", "Python", "Java", "C#", "Haskell"}
	fmt.Println(cursos3)

	// Maps, simple hash, dictionary
	currencies := make(map[string]string)

	currencies["MXN"] = "Mexican Peso"
	currencies["Peru"] = "Nuevo Sol"
	currencies["USD"] = "USD Dollar"

	fmt.Println(currencies["USD"])
	fmt.Println(currencies)

	// Deleting a Currency from Dictionary
	delete(currencies, "Peru")
	fmt.Println("After Deletion")
	fmt.Println(currencies)

	fmt.Println("Checking if an entry exists")
	// Way 1
	canada_currency := currencies["canada"]
	fmt.Println(canada_currency)

	// Way 2
	_, exist := currencies["Peru"]
	fmt.Println(exist)

	fmt.Println("===================================")
	fmt.Println("STRUCTURES IN GO")
	fmt.Println("===================================")

	// Way 1
	var ocurso Curso
	ocurso.id =1
	ocurso.nombre = "Luicito"
	ocurso.profesor = "Mariano"

	fmt.Println(ocurso)
	// Way 2
	ocurso2 := Curso{}
	ocurso2.id =2
	ocurso2.nombre = "Panchito"
	ocurso2.profesor = "Mariano"

	fmt.Println(ocurso2)

	// Way 3
	ocurso3 := Curso {
		id:3,
		nombre: "Basico3",
		profesor: "Juanga"}

	fmt.Println(ocurso3)
}