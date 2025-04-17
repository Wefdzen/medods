package service

import (
	"crypto/sha256"

	"golang.org/x/crypto/bcrypt"
)

func HashString(refToken string) (string, error) {
	// Делема надеюсь вы это читаете, в задание написанно дословно
	// исключительно в bcrypt, но у bcrypt 72 byte limit
	// я хотел argon2 использовать, но опять же в задание
	// написано исключительно и по итогу либа как я сделал
	// либо можно было поделить на части 72 байта и сложить, но
	// как по мне это тупо.
	hash := sha256.Sum256([]byte(refToken))
	bytes, err := bcrypt.GenerateFromPassword(hash[:], 14)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}
