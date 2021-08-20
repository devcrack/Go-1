package main

import "fmt"

func main()  {
	calification := 11
	/*
	if calification > 10 || calification < 0 {
		fmt.Println("Your calification is out of range")
	} else {
		if calification < 5 {
			fmt.Println("Not Approved")
		} else {
			fmt.Println("Approved")
		}
	}
	 */

	sex := "Male"
	if sex == "Male" && calification > 5{
		fmt.Println("This guy is a Male and does not a Donkey")
	} else {
		fmt.Println("This guy is a Transformer and and probably is a Donkey")
	}

	// Loops in Go

	//Classic For
	for i := 1; i <= 20; i++ {
		fmt.Println(i)
	}

	// Truncate a For
	for i := 1; i <= 20; i++ {
		if i == 10{
			break
		}
		fmt.Println(i)

	}
	// Using CONTINUE like in python
	// Truncate a For
	fmt.Println("Using CONTINUE like in Python")
	for i := 1; i <= 20; i++ {
		if i == 10{
			continue
		}
		fmt.Println(i)

	}

	fmt.Println("===================")
	fmt.Println("Using FOR RANGE")

	random_string := "Learning Go"
	for i, value := range random_string {
		fmt.Println(i, string(value))
	}

	//Skipping the INDEX just use _
	for _, value := range random_string {
		fmt.Println(string(value))
	}

}
