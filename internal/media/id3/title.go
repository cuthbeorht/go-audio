package id3

import "errors"

func Title(frames []Frame) (string, error) {

	for _, frame := range frames {
		if frame.Id == "TIT2" {
			return frame.Data, nil
		}
	}

	return "", errors.New("no title found")
}
