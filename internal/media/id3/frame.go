package id3

type Frame struct {
	Id    string
	Size  int
	Flags FrameFlags
	Data  string
}

type FrameFlags struct {
	IsLatin1 bool
	IsBOM    bool
}

func NewFrame(data []byte) Frame {
	frame := data[10:22]

	id := frameId(frame)
	size := frameSize(frame)
	flags := frameFlags(frame)
	frameData := frameData(data[21:], size)

	return Frame{Id: id, Size: size, Flags: flags, Data: frameData}
}

func frameId(rawHeader []byte) string {
	return string(rawHeader[:4])
}

func frameSize(data []byte) int {

	// size := data[4:8]
	var size int
	for _, value := range data[4:8] {
		size += int(value)
	}

	return size
}

func frameFlags(data []byte) FrameFlags {
	// rawFlags := data[9:11]
	isLatin1 := false
	isBOM := false

	// TODO: Don't understand what the flags should be
	// if rawFlags[0] == byte(0) {
	// 	isLatin1 = true
	// }

	// if rawFlags[1] == byte(1) {
	// 	isBOM = true
	// }

	return FrameFlags{IsLatin1: isLatin1, IsBOM: isBOM}
}

func frameData(data []byte, length int) string {
	return string(data[:length-2])
}
