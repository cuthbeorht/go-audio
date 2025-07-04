package main

import (
	"fmt"

	"github.com/cuthbeorht/go-audio/internal/command"
	"github.com/cuthbeorht/go-audio/internal/errorHandling"
	"github.com/cuthbeorht/go-audio/internal/media/files"
	"github.com/cuthbeorht/go-audio/internal/media/id3"
)

func main() {
	mediaFile, err := command.GetMediaFileNameFromArgs()
	errorHandling.Check(err, "Invalid file")

	mediaData, err := files.LoadFile(mediaFile)
	errorHandling.Check(err, "Unable to load file")

	isId3Present := id3.IsId3Present(mediaData)

	if isId3Present {
		id3Size := id3.Id3Size(mediaData)
		fmt.Printf("\n\nID3 tag size: %d\n", id3Size)

		id3Version, err := id3.Version(mediaData)
		errorHandling.Check(err, "Invalid verion")
		fmt.Printf("\nVersion: %s", id3Version)

		tagData := id3.Id3Tag(mediaData, id3Size)
		fmt.Printf("\nTag data: %s", tagData)
	}
}
