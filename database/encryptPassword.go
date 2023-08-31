package database

import (
	"errors"
	"fmt"
	"log"

	"github.com/amren1254/gin-docker/model"

	"golang.org/x/crypto/bcrypt"
)

func encrypt(password string) ([]byte, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, errors.New("failed to encrypt password, please provide another password")
	}
	fmt.Println("hash to store ", string(hash))
	//store this hash into database
	return hash, nil
}

func VerifyUserCredential(user model.UserDetails) (bool, error) {
	// retrive password hash from database
	var dbRepo DatabaseRepository
	retrievedUser, err := dbRepo.retrieveUserDetails(user)
	if err != nil {
		log.Println(err)
	}
	//compare hash from database
	if err := bcrypt.CompareHashAndPassword([]byte(retrievedUser.Password), []byte(user.Password)); err != nil {
		return false, err
	}
	return true, nil
}
