package id3

type FrameId struct {
	Id          string
	Description string
}

func FrameIds() []FrameId {
	frames := []FrameId{
		{Id: "COMM", Description: "Comment"},
		{Id: "TIT2", Description: "Title"},
	}

	return frames
}
