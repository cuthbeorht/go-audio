package command

import (
	"errors"
	"fmt"
	"os"
)

func GetMediaFileNameFromArgs() (fileName string, err error) {
	if len(os.Args) == 2 {
		return os.Args[1], nil
	}

	msg := fmt.Sprintf("Invalid number of parameters. Expected 2, got %d", len(os.Args))
	err = errors.New(msg)

	return "", err
}
