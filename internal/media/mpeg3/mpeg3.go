package mpeg3

import (
	"errors"
	"fmt"
)

const maxSize = 268435455

func IsId3Present(data []byte) bool {

	id3 := string(data[0:3])

	fmt.Printf("ID3 contents: %s", id3)
	return id3 == "ID3"
}

func Id3Size(data []byte) int {

	byteSizes := []int{0, 0, 128}
	size := 0
	for index, value := range data[6:10] {
		size = size << 8
		if index < 3 {
			size += int(value) * byteSizes[index]
		} else {
			size += int(value)
		}
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
