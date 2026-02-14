package controllers

import (
	"errors"
	"fmt"
	"os"

	"github.com/Arif14377/koda-b6-golang5/internal/views"
)

func StartApp() {
	for true {
		views.Home()
		inputMenu := views.InputAngka("\nInput menu: ")

		if inputMenu > 3 || inputMenu < 0 {
			err := errors.New("Pilihan tidak valid")
			fmt.Println(err)
			enterKosong := views.InputEnter()
			if enterKosong == "" {
				continue
			}
			fmt.Println("Input tidak valid")
			os.Exit(1)
		}

		switch inputMenu {
		case 1:
			RegisterController()
		case 2:
			LoginController()
		case 3:
			ForgotPasswordController()
		case 0:
			os.Exit(0)
		}
		break
	}
}

func RegisterController() {
	views.Register()
}

func LoginController() {
	views.Login()
}

func ForgotPasswordController() {
	views.ForgotPassword()
}
