package main

import "fmt"

func main()  {
	/*func() {
		fmt.Println("I am an anonymous function")
	}*/

	/*say_hi := func() {
		fmt.Println("I am other anonymous function")
	}
	say_hi()*/

	course_names("Javascript", "C #", "Go")

	fmt.Println("Pointers")
	a := 700
	triple_of_value(&a)
	fmt.Println(a)
}

func course_names(courses ...string)  {
	for _, v := range courses {
		fmt.Println(v)
	}
}

func triple_of_value(number *int) {
	*number = *number * 3

	fmt.Println("Address direction in memory ", number)
	fmt.Println("Varable value ", *number)
}
