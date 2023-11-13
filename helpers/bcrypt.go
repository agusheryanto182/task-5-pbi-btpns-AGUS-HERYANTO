package helpers

import "golang.org/x/crypto/bcrypt"

func HashPassword(pass string) string {
	hashResult, err := bcrypt.GenerateFromPassword([]byte(pass), bcrypt.MinCost)
	if err != nil {
		return err.Error()
	}

	return string(hashResult)
}

func ComparePassword(dbPass, inputPass string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(dbPass), []byte(inputPass))
	return err == nil
}
