package main

import (
	"fmt"
)

func main ()  {
	fmt.Println(operation(12, 14, "borrow"))
	fmt.Println(operation_switch(12, 14, "multiplication"))
	fmt.Println("Returning multiples values")
	suma, resta, multiplicacion := operations_return1(12, 23)
	fmt.Println("Addition ", suma, "  Borrow", resta, " multtiplication ", multiplicacion)
	suma2, resta2, multiplicacion2 := operations_noreturn(23, 43)
	fmt.Println("Addition ", suma2, "  Borrow", resta2, " multtiplication ", multiplicacion2)
}

func operation(num1, num2 int, type_calc string)  int {
	if type_calc == "add" {
		return num1 + num2
	 } else if type_calc == "borrow" {
		return num1 - num2
	} else if type_calc == "multiplication" {
		return num1 * num2
	}
	return 0
}

func operation_switch(num1, num2 int, type_calc string) int {

	switch {
	case type_calc == "add": return num1 + num2
	case type_calc == "borrow": return num1 - num2
	case type_calc == "multiplication": return num1 * num2
	default:
		return 0
	}
}

func operations_return1(num1, num2 int) (int, int, int){
	addition := num1 + num2
	borrow := num1 - num2
	multiplication := num1 * num2
	return addition, borrow, multiplication
}

func operations_noreturn(num1, num2 int)  (addition int, borrow int, mult int){
	addition = num1 + num2
	borrow = num1 - num2
	mult = num1 * num2
	return
}