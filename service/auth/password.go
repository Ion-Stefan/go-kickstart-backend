package auth

import "golang.org/x/crypto/bcrypt"

func HashPassword(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	return string(hash), nil
}

func ComparePasswords(hashed []byte, plain string) bool {
	// CompareHashAndPassword compares a bcrypt hashed password with its possible plaintext equivalent.
	// Returns nil on success, or an error on failure.
	err := bcrypt.CompareHashAndPassword(hashed, []byte(plain))
	return err == nil
}
