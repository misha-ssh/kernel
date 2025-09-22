package connect

import (
	"errors"
	"regexp"
	"strings"
)

var (
	errEmptyAlias        = errors.New("alias is empty")
	errSpecialCharacters = errors.New("special characters are not allowed")

	aliasPattern = regexp.MustCompile(`^[a-zA-Z0-9_]+$`)
)

func (c *Connect) Validate() error {
	return validateAlias(c.Alias)
}

func validateAlias(alias string) error {
	if strings.TrimSpace(alias) == "" {
		return errEmptyAlias
	}

	if !aliasPattern.MatchString(alias) {
		return errSpecialCharacters
	}

	return nil
}
