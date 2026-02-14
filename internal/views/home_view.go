package views

import "fmt"

func Home() {
	fmt.Println("--- Welcome ---")
	fmt.Println("1. Register")
	fmt.Println("2. Login")
	fmt.Println("3. Forgot Password")
	fmt.Println("\n0. Exit")
}
