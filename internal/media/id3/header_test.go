package id3

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"os/exec"
	"testing"
)

func getActualID3TagLength(path string) int {
	var output bytes.Buffer
	complexCmd := "ffmpeg -i sample.mp3 -v debug 2>&1 | grep id3v2 | awk '{print $4}' | awk -F ':' '{print $2}'"
	cmd := exec.Command("bash", "-c", complexCmd)
	cmd.Stdout = &output

	err := cmd.Run()
	if err != nil {
		log.Fatalf("Unable to run command: %s", err)
	}

	fmt.Printf("\nActual length: %s", output.String())

	return 121
}

func Test(t *testing.T) {
	testFileData, err := os.ReadFile("../../../sample.mp3")
	if err != nil {
		fmt.Println(os.Getwd())
		log.Fatalf("Unale to open sample mp3 file: %s", err)
	}
	expectedSize := getActualID3TagLength("fake")
	actualSize := Id3Size(testFileData)

	if expectedSize != actualSize {
		t.Errorf("Expected %d; Got %d", expectedSize, &actualSize)
	}
}