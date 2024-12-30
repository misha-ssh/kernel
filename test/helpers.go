package test

import (
	"fmt"
	"math/rand"
	"os/user"
)

const DefaultChars = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
const DefaultCountChars = 5

func RandomString() string {
	return generateRandomString(DefaultCountChars)
}

func GetDirForTests() string {
	usr, err := user.Current()

	if err != nil {
		fmt.Println(err)
	}

	return usr.HomeDir + "/test/"
}

func generateRandomString(count int) string {
	letterRunes := []rune(DefaultChars)

	randomStr := make([]rune, count)
	for i := range randomStr {
		randomStr[i] = letterRunes[rand.Intn(len(letterRunes))]
	}

	return string(randomStr)
}
