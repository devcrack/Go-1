package main
import "fmt"

func main()  {
	say_hi("Roger", "Perez", 32)
	say_hi("Juancho", "Dominguez", 21)
}

func say_hi(name, last_name string, age int)  {
	fmt.Println("Hello my name is", name, last_name, "and my age is: " , age, " years old")

}

