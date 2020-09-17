package framework

import (
	"strings"
	"errors"
)

func OneWordChecker(word string) error {
	if len(strings.Split(word, " ")) != 1 {
		return errors.New(word + ` must be one word`)
	}
	return nil
}