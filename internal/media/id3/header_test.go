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

func TestGivenValidID3TagID3SizeExpectValidSize(t *testing.T) {

	expectedSize := getActualID3TagLength("fake")
	actualSize := Size(GetSample())

	if expectedSize != actualSize {
		t.Errorf("Expected %d; Got %d", expectedSize, &actualSize)
	}
}

func TestGivenValidMP3VersionExpect2_4(t *testing.T) {
	sampleMediaFile := GetSample()

	expectedVersion := "2.4"
	actualVersion, _ := Version(sampleMediaFile)

	if expectedVersion != actualVersion {
		t.Errorf("Expected '%s'. Got '%s'", expectedVersion, actualVersion)
	}
}

func GetSample() []byte {
	testFileData, err := os.ReadFile("../../../sample.mp3")
	if err != nil {
		fmt.Println(os.Getwd())
		msg := fmt.Sprintf("Unale to open sample mp3 file: %s", err)
		log.Fatal(msg)
		panic(msg)
	}

	return testFileData
}
