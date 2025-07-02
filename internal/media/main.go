package media

import (
	"errors"
	"fmt"
)

func IsId3Present(data []byte) bool {

	id3 := string(data[0:3])

	fmt.Printf("ID3 contents: %s", id3)
	return id3 == "ID3"
}

func Id3Size(data []byte) int {

	var size int
	for _, value := range data[8:11] {
		size += int(value)
	}

	return size
}

func Version(data []byte) (string, error) {
	major := data[3]

	switch major {
	case 3:
	case 4:
		return "2.3", nil

	}

	return "", errors.New("missing ID3 version")
}

func Id3Tag(data []byte, size int) string {
	endByte := 12 + size
	rawTags := data[12:endByte]

	return string(rawTags)
}
