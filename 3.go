package main

import "fmt"

func main()  {
	//Declaring an Array
	var cursos[3] string
	cursos[0] = "Go"
	cursos[1] ="JavaScript"
	cursos[2] = "C sharp"
	for _, value := range cursos{
		fmt.Println(value)
	}
}