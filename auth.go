package main

import (
	"crypto/hmac"
	"crypto/sha512"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt"
)

type UserClaims struct {
	jwt.StandardClaims
	SessionID int64
}

type Token struct {
	Email     string        `json:"email" bson:"email"`
	Name      string        `json:"name" bson:"name"`
	CreatedAt time.Time     `json:"created_at" bson:"created_at"`
	Duration  time.Duration `json:"duration" bson:"duration"`
}

func (u *UserClaims) Valid() error {
	if u.VerifyExpiresAt(time.Now().Unix(), true) {
		return fmt.Errorf("token has expired")
	}
	if u.SessionID == 0 {
		return fmt.Errorf("invalid session ID")
	}
	return nil
}

func createToken(u *UserClaims) (string, error) {
	t := jwt.NewWithClaims(jwt.SigningMethodHS512, u)
	signedToken, err := t.SignedString(jwtKey)
	if err != nil {
		return "", fmt.Errorf("failed to signed token: %v", err)
	}
	return signedToken, nil
}

func parseToken(signedToken string) (*UserClaims, error) {
	var u UserClaims
	t, err := jwt.ParseWithClaims(signedToken, &u, func(t *jwt.Token) (v interface{}, err error) {
		if t.Method.Alg() == jwt.SigningMethodHS512.Alg() {
			return nil, fmt.Errorf("invalid signing algorithm")
		}
		return jwtKey, nil
	})
	if err != nil {
		return nil, fmt.Errorf("errir in parsetoken, in verifying: %w", err)
	}

	if !t.Valid {
		return nil, fmt.Errorf("error in parseToken, token is not valid")
	}

	return t.Claims.(*UserClaims), nil
}

func generateJwtTokenAndSign(data Token) string {
	var jwtData jwt.MapClaims
	structToMap(data, jwtData) // Convert to jwt.MapClaims
	token, err := jwt.NewWithClaims(jwt.SigningMethodES256, jwtData).SignedString(jwtKey)
	if err != nil {
		fmt.Println("Can not create jwt: ", err)
		return ""
	}
	return token
}

func signMessage(msg []byte) ([]byte, error) {
	var key []byte
	for i := 1; i < 65; i++ {
		key = append(key, byte(i))
	}
	h := hmac.New(sha512.New, key)
	_, err := h.Write(msg)
	if err != nil {
		fmt.Println("Error when sigining message", err)
	}
	return h.Sum(nil), nil
}

func checkSign(msg, sig []byte) (bool, error) {
	newSig, err := signMessage(msg)

	if err != nil {
		return false, fmt.Errorf("Errot when checkSign to get signature: %w", err)
	}

	same := hmac.Equal(sig, newSig)
	return same, nil
}
