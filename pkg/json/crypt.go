package json

import (
	"errors"

	crypt2 "github.com/ssh-connection-manager/kernel/v2/internal/encryption"

	"github.com/ssh-connection-manager/kernel/v2/internal/logger"
)

func SetCryptData(c Connect) (Connect, error) {
	var err error

	errMess := errors.New("err set crypt data")

	c.Alias, err = crypt2.Encrypt(c.Alias)
	if err != nil {
		logger.Danger(err.Error())
		return c, errMess
	}
	c.Address, err = crypt2.Encrypt(c.Address)
	if err != nil {
		logger.Danger(err.Error())
		return c, errMess
	}
	c.Login, err = crypt2.Encrypt(c.Login)
	if err != nil {
		logger.Danger(err.Error())
		return c, errMess
	}
	c.Password, err = crypt2.Encrypt(c.Password)
	if err != nil {
		logger.Danger(err.Error())
		return c, errMess
	}
	c.CreatedAt, err = crypt2.Encrypt(c.CreatedAt)
	if err != nil {
		logger.Danger(err.Error())
		return c, errMess
	}
	c.UpdatedAt, err = crypt2.Encrypt(c.UpdatedAt)
	if err != nil {
		logger.Danger(err.Error())
		return c, errMess
	}

	return c, nil
}

func decryptData(c Connect) (Connect, error) {
	var err error

	errMess := errors.New("err decrypt data")

	c.Alias, err = crypt2.Decrypt(c.Alias)
	if err != nil {
		logger.Danger(err.Error())
		return c, errMess
	}
	c.Address, err = crypt2.Decrypt(c.Address)
	if err != nil {
		logger.Danger(err.Error())
		return c, errMess
	}
	c.Login, err = crypt2.Decrypt(c.Login)
	if err != nil {
		logger.Danger(err.Error())
		return c, errMess
	}
	c.Password, err = crypt2.Decrypt(c.Password)
	if err != nil {
		logger.Danger(err.Error())
		return c, errMess
	}
	c.CreatedAt, err = crypt2.Decrypt(c.CreatedAt)
	if err != nil {
		logger.Danger(err.Error())
		return c, errMess
	}
	c.UpdatedAt, err = crypt2.Decrypt(c.UpdatedAt)
	if err != nil {
		logger.Danger(err.Error())
		return c, errMess
	}

	return c, nil
}

func (c *Connections) SetDecryptData() error {
	errMess := errors.New("err set decrypt data")

	for key, connect := range c.Connects {
		data, err := decryptData(connect)
		if err != nil {
			logger.Danger(err.Error())
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
			logger.Danger(err.Error())
			return errMess
		}

		c.Connects[key] = data
	}

	return nil
}
