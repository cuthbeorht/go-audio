package id3

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"strings"
	"testing"
)

func getActualID3TagLength(path string) int {

	ex, _ := os.Executable()
	fmt.Println("CWDL ", filepath.Dir(ex))

	cmd := exec.Command("/Users/davidsciacchettano/src/go-audio/scripts/audio-metadata.sh")
	stdout, stderr := cmd.Output()

	if stderr != nil {
		fmt.Printf("Error getting length: %s", stderr)
		panic("cannot get length")
	}
	num, _ := strconv.Atoi(strings.TrimSpace(string(stdout)))
	fmt.Printf("Length: %d", num)

	return num
}

func TestGivenValidMp3IsTagPresentExpectID3(t *testing.T) {
	expectedTagPresent := true
	actualTagPresent := IsTagPresent(GetSample())

	if expectedTagPresent != actualTagPresent {
		t.Errorf("Expected %t. Got %t", expectedTagPresent, actualTagPresent)
	}
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

	expectedVersion := "2.3"
	actualVersion, _ := Version(sampleMediaFile)

	if expectedVersion != actualVersion {
		t.Errorf("Expected '%s'. Got '%s'", expectedVersion, actualVersion)
	}
}

func GetSample() []byte {
	testFileData, err := os.ReadFile("../../../fixture-sample.dat")
	if err != nil {
		fmt.Println(os.Getwd())
		msg := fmt.Sprintf("Unale to open sample mp3 file: %s", err)
		log.Fatal(msg)
		panic(msg)
	}

	return testFileData
}
