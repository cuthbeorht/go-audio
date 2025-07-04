package id3

import "testing"

func TestGivernValidSampleNewFrameExpectTIT2Id(t *testing.T) {
	sampleMp3 := GetSample()

	expectedFrame := Frame{Id: "TIT2"}
	actualFrame := NewFrame(sampleMp3)

	if expectedFrame.Id != actualFrame.Id {
		t.Errorf("Expected '%s'. Got '%s'", expectedFrame.Id, actualFrame.Id)
	}
}

func TestGivernValidSampleNewFrameExpectTIT2Size(t *testing.T) {
	sampleMp3 := GetSample()

	expectedFrame := Frame{Size: 28}
	actualFrame := NewFrame(sampleMp3)

	if expectedFrame.Size != actualFrame.Size {
		t.Errorf("Expected '%s'. Got '%s'", expectedFrame.Id, actualFrame.Id)
	}
}
