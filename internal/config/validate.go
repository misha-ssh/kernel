package config

import (
	"regexp"
	"unicode"
)

func validateOnEmptyString(value string) error {
	if value == "" {
		return ErrValueIsInvalid
	}

	runes := []rune(value)
	for _, r := range runes {
		if unicode.IsSpace(r) {
			return ErrValueIsInvalid
		}
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
	matchedValue, err := regexp.MatchString("^[a-zA-Z0-9.\\-_]+$", value)
	if err != nil {
		return err
	}
	if !matchedValue {
		return ErrKeyOfNonLetters
	}

	err = validateOnEmptyString(value)
	if err != nil {
		return err
	}

	return nil
}
