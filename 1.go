package main
// Importing simple WAY
//import "fmt"
//import "strconv"

// Import refactored way
import (
	"fmt"
	"strconv"
)

func main() {
	// Declaring a variable
	var my_var string
	my_var = "hello boys"
	fmt.Println(my_var)
	// Declaring and assign a value at same time
	my_var2 := "Hello from Here vato"
	fmt.Println(my_var2)

	// Pythonic way of declare varibales and assign its values
	name, last_name, age := "Jhon", "Litt", 50
	fmt.Println(name, last_name, age)

	// Formatting prints
	name2:= "George"
	age2 := 50
	is_married := false
	// Very important use Printf for use formatting
	fmt.Printf("My Name is %s and my age is %d, married civil status: %t", name2, age2, is_married)

	// Casting variables
	num1 := 20
	num2 := "43"
	num3, _ := strconv.Atoi(num2)
	addition := num1 + num3
	fmt.Printf("\n%d plus %s is equal to:%d", num1, num2, addition)
}


