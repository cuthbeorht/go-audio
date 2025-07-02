package main

import (
	"fmt"

	"github.com/cuthbeorht/go-audio/internal/command"
	"github.com/cuthbeorht/go-audio/internal/errorHandling"
	"github.com/cuthbeorht/go-audio/internal/media"
)

func main() {
	mediaFile, err := command.GetMediaFileNameFromArgs()
	errorHandling.Check(err, "Invalid file")

	mediaData, err := media.LoadFile(mediaFile)
	errorHandling.Check(err, "Unable to load file")

	isId3Present := media.IsId3Present(mediaData)

	if isId3Present {
		id3Size := media.Id3Size(mediaData)
		fmt.Printf("\n\nID3 tag size: %d\n", id3Size)

		id3Version, err := media.Version(mediaData)
		errorHandling.Check(err, "Invalid verion")
		fmt.Printf("\nVersion: %s", id3Version)

		tagData := media.Id3Tag(mediaData, id3Size)
		fmt.Printf("\nTag data: %s", tagData)
	}
}
