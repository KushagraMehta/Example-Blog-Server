package util

import (
	"errors"
	"strings"
)

func FormatError(err error) error {

	if strings.Contains(err.Error(), "no rows in result set") {
		return errors.New("record don't exist")
	}
	return errors.New("Incorrect Details")
}
