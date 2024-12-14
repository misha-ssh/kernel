package test

import (
	"fmt"
	"math/rand"
	"os/user"
)

func RandomString(count int) string {
	letterRunes := []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

	randomStr := make([]rune, count)
	for i := range randomStr {
		randomStr[i] = letterRunes[rand.Intn(len(letterRunes))]
	}

	return string(randomStr)
}

func GetDirForTests() string {
	usr, err := user.Current()

	if err != nil {
		fmt.Println(err)
	}

	return usr.HomeDir + "/test/"
}
