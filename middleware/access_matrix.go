package middleware

import (
	"os"

	"gopkg.in/yaml.v3"
)

type AccessConfig map[string]map[string][]string

var AccessMatrix AccessConfig

func LoadAccessMatrix(path string) error {
	data, err := os.ReadFile(path)
	if err != nil {
		return err
	}

	err = yaml.Unmarshal(data, &AccessMatrix)
	if err != nil {
		return err
	}

	return nil
}
