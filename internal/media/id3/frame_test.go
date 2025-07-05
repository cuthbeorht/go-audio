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

func TestGivernValidSampleNewFrameExpectValidFlags(t *testing.T) {
	sampleMp3 := GetSample()

	expectedFrame := Frame{Flags: FrameFlags{IsLatin1: false}}
	actualFrame := NewFrame(sampleMp3)

	if expectedFrame.Flags.IsLatin1 != actualFrame.Flags.IsLatin1 {
		t.Errorf("Expected '%s'. Got '%s'", expectedFrame.Id, actualFrame.Id)
	}
}

func TestGivernValidSampleNewFrameExpectValidData(t *testing.T) {
	sampleMp3 := GetSample()

	expectedFrame := Frame{Data: "Quiet Saturday 001 (00:17)"}
	actualFrame := NewFrame(sampleMp3)

	if expectedFrame.Data != actualFrame.Data {
		t.Errorf("Expected '%s'. Got '%s'", expectedFrame.Data, actualFrame.Data)
	}
}
