package services

import (
	"errors"
	"fmt"
	"strings"

	"github.com/Arif14377/koda-b6-golang5/internal/models"
)

var users []models.User

func Register(user models.User, confirmPassword string) error {
	// input tidak boleh kosong
	if user.Email == "" || user.FirstName == "" || user.LastName == "" || user.Password == "" {
		err := errors.New("Input tidak boleh kosong.")
		return err
	}

	// Email yang sudah terdaftar tidak bisa digunakan.
	for i := range len(users) {
		if users[i].Email == user.Email {
			err := errors.New("Email sudah terdaftar.")
			return err
		}
	}

	// Email tidak valid
	if !strings.Contains(user.Email, "@") {
		err := errors.New("Email tidak valid.")
		return err
	}

	// Password harus lebih dari 6 karakter
	if len(user.Password) < 6 {
		err := errors.New("Password minimal 6 karakter.")
		return err
	}

	// Confirm password harus sama dengan password
	if confirmPassword != user.Password {
		err := errors.New("Password tidak sama.")
		return err
	}

	users = append(users, user)
	fmt.Println(users)
	return nil
}
