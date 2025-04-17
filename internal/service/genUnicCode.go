package service

import (
	"fmt"
	"math/rand"
)

// GenUnicCode generate random code for couple of tokens.
func GenUnicCode() string {
	return fmt.Sprintf("%v", rand.Intn(10000))
}
