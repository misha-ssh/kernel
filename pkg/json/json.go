package json

import (
	"encoding/json"
	"errors"
)

func (c *Connections) SerializationJson(dataConnectsInFile string) error {
	err := json.Unmarshal([]byte(dataConnectsInFile), &c)

	if err != nil {
		return errors.New("error serializing json: " + err.Error())
	}

	return nil
}

func (c *Connections) deserializationJson() ([]byte, error) {
	newDataConnect, err := json.MarshalIndent(&c, "", " ")

	if err != nil {
		return nil, errors.New("error deserializing json: " + err.Error())
	}

	return newDataConnect, nil
}
