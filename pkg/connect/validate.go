package connect

import (
	"encoding/pem"
	"errors"
	"golang.org/x/crypto/ssh"
	"net"
	"os"
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
		c.validateAlias(),
		c.validatePrivateKey(),
		c.validatePassword(),
		c.validateLogin(),
		c.validateAddress(),
		c.validateCreatedAt(),
		c.validateUpdatedAt(),
		c.validatePort(),
	} {
		if err != nil {
			return err
		}
	}
	return nil
}

func (c *Connect) validateAlias() error {
	if strings.TrimSpace(c.Alias) == "" {
		return errors.New("alias is empty")
	}

	if !aliasPattern.MatchString(c.Alias) {
		return errors.New("alias special characters are not allowed")
	}

	return nil
}

func (c *Connect) validateLogin() error {
	if strings.TrimSpace(c.Login) == "" {
		return errors.New("login cannot be empty")
	}

	if len(c.Login) > 50 {
		return errors.New("login too long (max 50 characters)")
	}

	if !loginPattern.MatchString(c.Login) {
		return errors.New("login contains invalid characters")
	}

	return nil
}

func (c *Connect) validateAddress() error {
	if strings.TrimSpace(c.Address) == "" {
		return errors.New("address cannot be empty")
	}

	if ip := net.ParseIP(c.Address); ip != nil {
		return nil
	}

	if !addressPattern.MatchString(c.Address) {
		return errors.New("invalid address format")
	}

	if len(c.Address) > 253 {
		return errors.New("address too long")
	}

	return nil
}

func (c *Connect) validatePassword() error {
	if strings.TrimSpace(c.PrivateKey) != "" {
		return nil
	}

	if strings.TrimSpace(c.Password) == "" {
		return errors.New("password cannot be empty")
	}

	if len(c.Password) < 4 {
		return errors.New("password too short (min 4 characters)")
	}

	if len(c.Password) > 100 {
		return errors.New("password too long (max 100 characters)")
	}

	return nil
}

func (c *Connect) validatePrivateKey() error {
	if strings.TrimSpace(c.Password) != "" {
		return nil
	}

	data, err := os.ReadFile(c.PrivateKey)
	if err != nil {
		return errors.New("note found private key")
	}

	block, _ := pem.Decode(data)
	if block == nil {
		return errors.New("private key is not valid")
	}

	_, err = ssh.ParseRawPrivateKey(data)
	if err != nil {
		if !strings.Contains(err.Error(), "passphrase") {
			return err
		}

		_, err = ssh.ParsePrivateKeyWithPassphrase(data, []byte(c.Passphrase))
	}

	return err
}

func (c *Connect) validateCreatedAt() error {
	return validateDate(c.CreatedAt)
}

func (c *Connect) validateUpdatedAt() error {
	return validateDate(c.UpdatedAt)
}

func validateDate(date string) error {
	if strings.TrimSpace(date) == "" {
		return nil
	}

	parsedTime, err := time.Parse(time.RFC3339, date)
	if err != nil {
		return err
	}

	if parsedTime.After(time.Now()) {
		return errors.New("date cannot be in the future")
	}

	return nil
}

func (c *Connect) validatePort() error {
	if c.Port < 1 || c.Port > 65535 {
		return errors.New("port must be between 1 and 65535")
	}

	return nil
}
