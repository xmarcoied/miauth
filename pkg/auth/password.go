package auth

import "golang.org/x/crypto/bcrypt"

//GenerateHashPassword generates a password hash
func (s *AuthService) GenerateHashPassword(passowrd string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(passowrd), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}

//IsHashPasswordValid compares a password and a hash
func (s *AuthService) IsHashPasswordValid(hashedPassword, plainPassword string) (bool, error) {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(plainPassword))
	if err == nil {
		return true, nil
	}

	//Password does not match!
	if err == bcrypt.ErrMismatchedHashAndPassword {
		return false, nil
	}
	return false, err
}
