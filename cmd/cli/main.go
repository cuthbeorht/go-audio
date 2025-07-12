package main

import (
	"fmt"

	"github.com/cuthbeorht/id3-tag-editor/internal/command"
	"github.com/cuthbeorht/id3-tag-editor/internal/errorHandling"
	"github.com/cuthbeorht/id3-tag-editor/internal/media/files"
	"github.com/cuthbeorht/id3-tag-editor/internal/media/id3"
)

func main() {
	mediaFile, err := command.GetMediaFileNameFromArgs()
	errorHandling.Check(err, "Invalid file")

	mediaData, err := files.LoadFile(mediaFile)
	errorHandling.Check(err, "Unable to load file")

	isId3Present := id3.IsTagPresent(mediaData)

	if isId3Present {
		newTag := id3.NewTag(mediaData)
		fmt.Println("Frames: ", newTag.Frames)
	}
}
