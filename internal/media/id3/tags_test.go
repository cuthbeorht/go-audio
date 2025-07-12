package id3

import "testing"

func TestGivenValidSampleNewTagExpectMultipleFrames(t *testing.T) {
	sampleData := GetSample()
	expectedFrames := 4

	newTag := NewTag(sampleData)

	if len(newTag.Frames) != expectedFrames {
		t.Errorf("Expected '%d'. Got '%d", len(newTag.Frames), expectedFrames)
	}
}

func TestGivernValidSampleNewFrameExpectTIT2Id(t *testing.T) {
	sampleMp3 := GetSample()

	expectedFrame := Frame{Id: "TIT2"}
	actualFrame, _ := NewFrame(sampleMp3, 10)

	if expectedFrame.Id != actualFrame.Id {
		t.Errorf("Expected '%s'. Got '%s'", expectedFrame.Id, actualFrame.Id)
	}
}

func TestGivernValidSampleNewFrameExpectTIT2Size(t *testing.T) {
	sampleMp3 := GetSample()

	expectedFrame := Frame{Size: 28}
	actualFrame, _ := NewFrame(sampleMp3, 10)

	if expectedFrame.Size != actualFrame.Size {
		t.Errorf("Expected '%s'. Got '%s'", expectedFrame.Id, actualFrame.Id)
	}
}

func TestGivernValidSampleNewFrameExpectValidFlags(t *testing.T) {
	sampleMp3 := GetSample()

	expectedFrame := Frame{Flags: FrameFlags{IsLatin1: false}}
	actualFrame, _ := NewFrame(sampleMp3, 10)

	if expectedFrame.Flags.IsLatin1 != actualFrame.Flags.IsLatin1 {
		t.Errorf("Expected '%s'. Got '%s'", expectedFrame.Id, actualFrame.Id)
	}
}

func TestGivernValidSampleNewFrameExpectValidData(t *testing.T) {
	sampleMp3 := GetSample()

	expectedFrame := Frame{Data: "Quiet Saturday 001 (00:17)"}
	actualFrame, _ := NewFrame(sampleMp3, 10)

	if expectedFrame.Data != actualFrame.Data {
		t.Errorf("Expected '%s'. Got '%s'", expectedFrame.Data, actualFrame.Data)
	}
}
