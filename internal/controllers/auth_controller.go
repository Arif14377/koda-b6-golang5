package controllers

import (
	"errors"
	"fmt"
	"os"

	"github.com/Arif14377/koda-b6-golang5/internal/models"
	"github.com/Arif14377/koda-b6-golang5/internal/services"
	"github.com/Arif14377/koda-b6-golang5/internal/views"
)

func StartApp() {
	for true {
		views.Home()
		inputMenu := views.InputAngka("\nInput menu: ")

		if inputMenu > 3 || inputMenu < 0 {
			err := errors.New("\nPilihan tidak valid")
			fmt.Println(err)
			enterKosong := views.InputEnter()
			if enterKosong == "" {
				continue
			}
			fmt.Println("\nInput tidak valid")
			os.Exit(1)
		}

		// FIXME : belum bener. palingan caranya input pakai string dan validasi diconvert ke number.
		// if _, err := strconv.Atoi(string(inputMenu)); err == nil {
		// 	fmt.Println(errors.New("\nInput tidak valid."))
		// 	enterKosong := views.InputEnter()
		// 	if enterKosong == "" {
		// 		continue
		// 	}
		// 	fmt.Println("\nInput tidak valid")
		// 	os.Exit(1)
		// }

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
	for true {
		views.Register()

		firstName := views.InputString("What is your first name: ")
		lastName := views.InputString("What is your last name: ")
		email := views.InputString("What is your email: ")
		password := views.InputString("Enter a strong password: ")
		confirmPassword := views.InputString("Confirm your password: ")

		user := models.User{
			FirstName: firstName,
			LastName:  lastName,
			Email:     email,
			Password:  password,
		}

		err := services.Register(user, confirmPassword)

		if err != nil {
			fmt.Println(err)
			continue
		}

		fmt.Println("Registrasi berhasil.")

		LoginController()
	}
}

func LoginController() {
	for true {
		views.Login()
		email := views.InputString("Enter your email: ")
		password := views.InputString("Enter your password: ")

		user := models.User{
			Email:    email,
			Password: password,
		}

		err := services.Login(user)

		if err != nil {
			fmt.Println(err)
			continue
		}
		fmt.Println("Login berhasil berhasil.")
		os.Exit(0)
	}
}

func ForgotPasswordController() {
	views.ForgotPassword()
}
