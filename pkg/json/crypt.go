package json

import (
	"errors"

	"github.com/ssh-connection-manager/crypt"
)

func SetCryptData(c Connect) (Connect, error) {
	var err error

	errMess := errors.New("err set crypt data")

	c.Alias, err = crypt.Encrypt(c.Alias)
	if err != nil {
		return c, errMess
	}
	c.Address, err = crypt.Encrypt(c.Address)
	if err != nil {
		return c, errMess
	}
	c.Login, err = crypt.Encrypt(c.Login)
	if err != nil {
		return c, errMess
	}
	c.Password, err = crypt.Encrypt(c.Password)
	if err != nil {
		return c, errMess
	}
	c.CreatedAt, err = crypt.Encrypt(c.CreatedAt)
	if err != nil {
		return c, errMess
	}
	c.UpdatedAt, err = crypt.Encrypt(c.UpdatedAt)
	if err != nil {
		return c, errMess
	}

	return c, nil
}

func decryptData(c Connect) (Connect, error) {
	var err error

	errMess := errors.New("err decrypt data")

	c.Alias, err = crypt.Decrypt(c.Alias)
	if err != nil {
		return c, errMess
	}
	c.Address, err = crypt.Decrypt(c.Address)
	if err != nil {
		return c, errMess
	}
	c.Login, err = crypt.Decrypt(c.Login)
	if err != nil {
		return c, errMess
	}
	c.Password, err = crypt.Decrypt(c.Password)
	if err != nil {
		return c, errMess
	}
	c.CreatedAt, err = crypt.Decrypt(c.CreatedAt)
	if err != nil {
		return c, errMess
	}
	c.UpdatedAt, err = crypt.Decrypt(c.UpdatedAt)
	if err != nil {
		return c, errMess
	}

	return c, nil
}

func (c *Connections) SetDecryptData() error {
	errMess := errors.New("err set decrypt data")

	for key, connect := range c.Connects {
		data, err := decryptData(connect)
		if err != nil {
			return errMess
		}

		c.Connects[key] = data
	}

	return nil
}

func (c *Connections) SetCryptAllData() error {
	errMess := errors.New("err set decrypt data all")

	for key, connect := range c.Connects {
		data, err := SetCryptData(connect)
		if err != nil {
			return errMess
		}

		c.Connects[key] = data
	}

	return nil
}
