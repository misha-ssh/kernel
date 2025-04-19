package config

import (
	"regexp"
	"strings"
)

func validateOnEmptyString(value string) error {
	if len(strings.Fields(value)) > 1 || value == "" {
		return ErrValueIsInvalid
	}

	return nil
}

func validateKey(key string) error {
	matchedKey, err := regexp.MatchString("^[a-zA-Z]", key)
	if err != nil {
		return err
	}
	if !matchedKey {
		return ErrKeyOfNonLetters
	}

	err = validateOnEmptyString(key)
	if err != nil {
		return err
	}

	return nil
}

func validateValue(value string) error {
	err := validateOnEmptyString(value)
	if err != nil {
		return err
	}

	return nil
}
