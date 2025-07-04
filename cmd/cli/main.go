package main

import (
	"fmt"

	"github.com/cuthbeorht/go-audio/internal/command"
	"github.com/cuthbeorht/go-audio/internal/errorHandling"
	"github.com/cuthbeorht/go-audio/internal/media/files"
	"github.com/cuthbeorht/go-audio/internal/media/mpeg3"
)

func main() {
	mediaFile, err := command.GetMediaFileNameFromArgs()
	errorHandling.Check(err, "Invalid file")

	mediaData, err := files.LoadFile(mediaFile)
	errorHandling.Check(err, "Unable to load file")

	isId3Present := mpeg3.IsId3Present(mediaData)

	if isId3Present {
		id3Size := mpeg3.Id3Size(mediaData)
		fmt.Printf("\n\nID3 tag size: %d\n", id3Size)

		id3Version, err := mpeg3.Version(mediaData)
		errorHandling.Check(err, "Invalid verion")
		fmt.Printf("\nVersion: %s", id3Version)

		tagData := mpeg3.Id3Tag(mediaData, id3Size)
		fmt.Printf("\nTag data: %s", tagData)
	}
}
