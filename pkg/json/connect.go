package json

import (
	"errors"
)

type Connections struct {
	Connects []Connect `json:"connects"`
}

type Connect struct {
	Alias    string `json:"alias"`
	Login    string `json:"login"`
	Address  string `json:"address"`
	Password string `json:"password"`

	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

func (c *Connections) getConnectData() error {
	fileConn := GetFile()

	data, err := fileConn.ReadFile()
	if err != nil {
		return err
	}

	err = c.SerializationJson(data)
	if err != nil {
		return err
	}

	err = c.SetDecryptData()
	if err != nil {
		return err
	}

	return nil
}

func (c *Connections) GetDataForListConnect() ([][]string, error) {
	var result [][]string

	err := c.getConnectData()
	if err != nil {
		return result, err
	}

	for _, v := range c.Connects {
		newElement := []string{v.Alias, v.CreatedAt, v.UpdatedAt}
		result = append(result, newElement)
	}

	if len(result) == 0 {
		return result, errors.New("no connection found")
	}

	return result, nil
}

func (c *Connections) GetConnectionsAlias() ([]string, error) {
	var result []string

	err := c.getConnectData()
	if err != nil {
		return result, err
	}

	for _, conn := range c.Connects {
		result = append(result, conn.Alias)
	}

	if len(result) == 0 {
		return result, errors.New("no connection found")
	}

	return result, nil
}

func (c *Connections) ExistConnectJsonByIndex(alias string) (int, error) {
	var noFound = -1

	err := c.getConnectData()
	if err != nil {
		return noFound, err
	}

	defer func() {
		err = c.SetCryptAllData()
	}()

	for i, v := range c.Connects {
		if v.Alias == alias {
			return i, errors.New("the alias is already in use")
		}
	}

	return noFound, nil
}

func (c *Connections) WriteConnectToJson(connect Connect) error {
	_, err := c.ExistConnectJsonByIndex(connect.Alias)
	if err != nil {
		return err
	}

	fileConn := GetFile()

	data, err := fileConn.ReadFile()
	if err != nil {
		return err
	}

	err = c.SerializationJson(data)
	if err != nil {
		return err
	}

	encodedConnect, err := SetCryptData(connect)
	if err != nil {
		return err
	}

	c.Connects = append(c.Connects, encodedConnect)

	deserializationJson, err := c.deserializationJson()
	if err != nil {
		return err
	}

	err = fileConn.WriteFile(deserializationJson)
	if err != nil {
		return err
	}

	return nil
}

func (c *Connections) updateJsonDataByIndex(index int, connect Connect) error {
	if index >= 0 && index < len(c.Connects) {
		c.Connects[index].Alias = connect.Alias
		c.Connects[index].Address = connect.Address
		c.Connects[index].Login = connect.Login
		c.Connects[index].Password = connect.Password
		c.Connects[index].UpdatedAt = connect.UpdatedAt

		return nil
	}

	return errors.New("connection update error")
}

func (c *Connections) UpdateConnectJson(alias string, connect Connect) error {
	index, err := c.ExistConnectJsonByIndex(alias)
	if err != nil {
		return err
	}

	fileConn := GetFile()

	cryptData, err := SetCryptData(connect)
	if err != nil {
		return err
	}

	err = c.updateJsonDataByIndex(index, cryptData)
	if err != nil {
		return err
	}

	deserializationJson, err := c.deserializationJson()
	if err != nil {
		return err
	}

	err = fileConn.WriteFile(deserializationJson)
	if err != nil {
		return err
	}

	return nil
}

func (c *Connections) deleteJsonDataByIndex(index int) {
	copy(c.Connects[index:], c.Connects[index+1:])

	c.Connects[len(c.Connects)-1] = Connect{}
	c.Connects = c.Connects[:len(c.Connects)-1]
}

func (c *Connections) DeleteConnectToJson(alias string) error {
	index, err := c.ExistConnectJsonByIndex(alias)
	if err != nil {
		return err
	}

	c.deleteJsonDataByIndex(index)

	deserializationJson, err := c.deserializationJson()
	if err != nil {
		return err
	}

	fileConn := GetFile()
	err = fileConn.WriteFile(deserializationJson)
	if err != nil {
		return err
	}

	return nil
}
