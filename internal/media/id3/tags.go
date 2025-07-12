package id3

import (
	"errors"
	"strings"
)

type Frame struct {
	Id    string
	Size  int
	Flags FrameFlags
	Data  string
}

type FrameFlags struct {
	IsLatin1 bool
	IsBOM    bool
}

type Tag struct {
	Header Header
	Frames []Frame
}

func NewTag(data []byte) Tag {

	var frames []Frame
	var currentStartingByte int

	currentStartingByte = 0

	for {

		newFrame, err := NewFrame(data, currentStartingByte+10)

		if err != nil {
			break
		}

		frames = append(frames, newFrame)
		currentStartingByte += newFrame.Size + 10
	}

	return Tag{Frames: frames}
}

func NewFrame(data []byte, startingByte int) (Frame, error) {

	// frameDataStartingByte := startingByte + 21

	frame := data[startingByte:]

	id, err := frameId(frame)
	if err != nil {
		return Frame{}, errors.New("not a valid frame")
	}

	size := frameSize(frame)
	flags := frameFlags(frame)
	frameData := frameData(frame, size)

	return Frame{Id: id, Size: size, Flags: flags, Data: frameData}, nil
}

func frameId(rawHeader []byte) (string, error) {
	validPrefixes := []string{"T", "C"}

	id := string(rawHeader[:4])
	for _, prefix := range validPrefixes {
		if !strings.HasPrefix(id, "T") {
			return prefix, nil
		}
	}

	return "", errors.New("not a valid frame id")
}

func frameSize(data []byte) int {

	// size := data[4:8]
	var size int
	for _, value := range data[4:8] {
		size += int(value)
	}

	return size
}

func frameFlags(data []byte) FrameFlags {
	// rawFlags := data[9:11]
	isLatin1 := false
	isBOM := false

	// TODO: Don't understand what the flags should be
	// if rawFlags[0] == byte(0) {
	// 	isLatin1 = true
	// }

	// if rawFlags[1] == byte(1) {
	// 	isBOM = true
	// }

	return FrameFlags{IsLatin1: isLatin1, IsBOM: isBOM}
}

func frameData(data []byte, length int) string {
	startingByte := 11                      // The full data from the MP3 Header where the frame data starts after the header, byte 11
	endingByte := length + startingByte - 1 // The ending byte for the frame data is the starting byte with the length subtracting 2 bytes... Not sure why yet
	return string(data[startingByte:endingByte])
}
