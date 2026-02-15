package services

import (
	"crypto/md5"
	"encoding/hex"
	"errors"
	"fmt"
	"strings"

	"github.com/Arif14377/koda-b6-golang5/internal/models"
)

var users []models.User

func Register(user models.User, confirmPassword string) error {
	// input tidak boleh kosong
	if user.Email == "" || user.FirstName == "" || user.LastName == "" || user.Password == "" {
		err := errors.New("\nInput tidak boleh kosong.")
		return err
	}

	// Email yang sudah terdaftar tidak bisa digunakan.
	for i := range len(users) {
		if users[i].Email == user.Email {
			err := errors.New("\nEmail sudah terdaftar.")
			return err
		}
	}

	// Email tidak valid
	if !strings.Contains(user.Email, "@") {
		err := errors.New("\nEmail tidak valid.")
		return err
	}

	// Password harus lebih dari 6 karakter
	if len(user.Password) < 6 {
		err := errors.New("\nPassword minimal 6 karakter.")
		return err
	}

	// Confirm password harus sama dengan password
	if confirmPassword != user.Password {
		err := errors.New("\nPassword tidak sama.")
		return err
	}

	// Encrypt/hashing password dengan md5
	data := []byte(user.Password)
	hashPassword := md5.Sum(data)
	user.Password = hex.EncodeToString(hashPassword[:])
	users = append(users, user)
	fmt.Printf("\nNama Anda adalah: %s\n", user.FullName())
	fmt.Printf("Email Anda adalah : %s\n", user.Email)
	fmt.Printf("Password Anda (Encrypted): %s\n", user.Password)
	return nil
}
