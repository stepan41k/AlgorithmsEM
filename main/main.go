package main 

import (
	"fmt"
)

type User struct {
	age int
	firstname string
	lastname string
}

func main() {
	
	newUser := User{age: 10, firstname: "Petya", lastname: "Ivamov"}
	fmt.Println(newUser)

}