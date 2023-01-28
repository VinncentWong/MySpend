package util

import "golang.org/x/crypto/bcrypt"

func Hash(rawPassword string) (*[]byte, error) {
	newString, err := bcrypt.GenerateFromPassword([]byte(rawPassword), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	return &newString, nil
}

func CompareHash(rawPassword, hashPassword string) error {
	err := bcrypt.CompareHashAndPassword([]byte(hashPassword), []byte(rawPassword))
	return err
}
