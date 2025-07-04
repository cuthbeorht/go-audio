package id3

import "testing"

func TestGivenValidMP3TitleExpectTitle(t *testing.T) {
	expectedtitle := "fake"

	actualTitle, _ := Title(GetSample())

	if expectedtitle != actualTitle {
		t.Errorf("Expected %s. Got %s", expectedtitle, actualTitle)
	}
}