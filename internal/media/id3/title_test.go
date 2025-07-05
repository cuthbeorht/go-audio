package id3

import "testing"

func TestGivenValidMP3TitleExpectTitle(t *testing.T) {
	expectedtitle := "Quiet Saturday 024 (00:41)"

	actualTitle, _ := Title([]Frame{{Id: "TIT2", Data: "Quiet Saturday 024 (00:41)"}})

	if expectedtitle != actualTitle {
		t.Errorf("Expected '%s'. Got '%s'", expectedtitle, actualTitle)
	}
}
