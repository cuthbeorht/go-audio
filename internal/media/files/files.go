package files

import (
	"errors"
	"fmt"
	"os"
)

func LoadFile(path string) (data []byte, err error) {
	data, err = os.ReadFile(path)
	if err != nil {
		errMsg := fmt.Sprintf("Could not open file: %s", err)
		return nil, errors.New(errMsg)
	}

	return data, nil
}
