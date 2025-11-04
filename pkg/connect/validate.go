package connect

import (
	"errors"
	"net"
	"regexp"
	"strings"
	"time"
)

var (
	aliasPattern   = regexp.MustCompile(`^[a-zA-Z0-9_-]+$`)
	loginPattern   = regexp.MustCompile(`^[a-zA-Z0-9_.-]+$`)
	addressPattern = regexp.MustCompile(`^[a-zA-Z0-9.-]+$`)
)

func (c *Connect) Validate() error {
	for _, err := range []error{
		validateAlias(c.Alias),
		validatePassword(c.Password, c.SshOptions.PrivateKey),
		validateLogin(c.Login),
		validateAddress(c.Address),
		validateCreatedAt(c.CreatedAt),
		validateUpdatedAt(c.UpdatedAt),
		validatePort(c.SshOptions.Port),
	} {
		if err != nil {
			return err
		}
	}
	return nil
}

func validateAlias(alias string) error {
	if strings.TrimSpace(alias) == "" {
		return errors.New("alias is empty")
	}

	if !aliasPattern.MatchString(alias) {
		return errors.New("special characters are not allowed")
	}

	return nil
}

func validateLogin(login string) error {
	if strings.TrimSpace(login) == "" {
		return errors.New("login cannot be empty")
	}

	if len(login) > 50 {
		return errors.New("login too long (max 50 characters)")
	}

	if !loginPattern.MatchString(login) {
		return errors.New("login contains invalid characters")
	}

	return nil
}

func validateAddress(address string) error {
	if strings.TrimSpace(address) == "" {
		return errors.New("address cannot be empty")
	}

	if ip := net.ParseIP(address); ip != nil {
		return nil
	}

	if !addressPattern.MatchString(address) {
		return errors.New("invalid address format")
	}

	if len(address) > 253 {
		return errors.New("address too long")
	}

	return nil
}

func validatePassword(password string, privateKey string) error {
	if strings.TrimSpace(privateKey) != "" {
		return nil
	}

	if strings.TrimSpace(password) != password {
		return errors.New("password contains spaces")
	}

	if len(password) < 4 {
		return errors.New("password too short (min 4 characters)")
	}

	if len(password) > 100 {
		return errors.New("password too long (max 100 characters)")
	}

	return nil
}

func validateCreatedAt(date string) error {
	return validateDate(date)
}

func validateUpdatedAt(date string) error {
	return validateDate(date)
}

func validateDate(date string) error {
	if strings.TrimSpace(date) == "" {
		return errors.New("cannot be empty")
	}

	parsedTime, err := time.Parse(time.RFC3339, date)
	if err != nil {
		return err
	}

	if parsedTime.After(time.Now()) {
		return errors.New("cannot be in the future")
	}

	return nil
}

func validatePort(port int) error {
	if port < 1 || port > 65535 {
		return errors.New("port must be between 1 and 65535")
	}

	return nil
}
