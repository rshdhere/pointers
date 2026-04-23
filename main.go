package main

import "fmt"

type User struct {
	Name *string
}

func main() {
	name := "Raashed"
	u := User{Name: &name}
	fmt.Println(u.Name)

	if u.Name != nil {
		fmt.Println(*u.Name)
	}
}
