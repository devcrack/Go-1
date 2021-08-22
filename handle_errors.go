package main

import (
	"fmt"
	"errors"
)

func main()  {
	res, error := operation_switch(23, 5, "add")
	if error != nil {
		fmt.Println(error)
	} else {
		fmt.Println(res)
	}
}

func operation_switch(num1, num2 int, type_calc string) (result int, err error) {

	if num1 < 0 || num2 < 0 {
		err = errors.New("One of the number  is negative")
	}
	switch {
	case type_calc == "add": result = num1 + num2
	case type_calc == "borrow": result = num1 - num2
	case type_calc == "multiplication": result = num1 * num2
	default:
		result = 0
	}
	return
}
