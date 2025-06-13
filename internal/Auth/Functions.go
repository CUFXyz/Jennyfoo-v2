package auth

import (
	"fmt"
	"time"
	"v2/models"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

func (a *Authentificator) CryptPassword(password []byte) (string, error) {
	hashedPass, err := bcrypt.GenerateFromPassword(password, bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	return string(hashedPass), nil
}

func (a *Authentificator) AuthUser(hashedpassword string, nonhashedpassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedpassword), []byte(nonhashedpassword))
}

func (a *Authentificator) GenerateToken(user *models.User) string {
	payload := jwt.MapClaims{
		"sub": user.Login,
		"exp": time.Now().Add(time.Minute * 30).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
	signedToken, err := token.SignedString([]byte("DummySign"))
	if err != nil {
		fmt.Printf("Generate Token failed\n")
		return ""
	}
	return signedToken
}
