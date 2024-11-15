package app

import (
	"os"
	"strings"
)

func ReadUsernames(filename string) ([]string, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return strings.Fields(string(data)), nil
}
