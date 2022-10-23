package infra_helper

import (
	"log"
	"math/rand"
	"time"

	"github.com/dgrijalva/jwt-go"
)

var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

type DataForToken struct {
	Name  string
	Phone string
	Role  string
	Time  time.Time
	*jwt.StandardClaims
}

// func for generate random password with N string's long
func RandPass(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

// func for generate token
func GetToken(name, phone, role string) string {
	data := DataForToken{
		Name:  name,
		Phone: phone,
		Role:  role,
		Time:  time.Now(),
	}
	sign := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), data)
	token, _ := sign.SignedString([]byte("secret"))

	return token
}

func ExtractClaims(tokenStr string) (jwt.MapClaims, bool) {
	hmacSecretString := "secret"
	hmacSecret := []byte(hmacSecretString)
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		return hmacSecret, nil
	})

	if err != nil {
		return nil, false
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, true
	} else {
		log.Printf("Invalid JWT Token")
		return nil, false
	}
}
