package views

import "fmt"

func InputString(prompt string) string {
	var input string

	fmt.Printf("%s", prompt)
	fmt.Scanf("%s", &input)

	return input
}

func InputAngka(prompt string) int {
	var input int

	fmt.Printf("%s", prompt)
	fmt.Scanf("%d", &input)

	return input
}

func InputEnter() string {
	var enterKosong string
	fmt.Printf("%s", "Enter untuk melanjutkan..")
	fmt.Scanf("%s", &enterKosong)
	return enterKosong
}
