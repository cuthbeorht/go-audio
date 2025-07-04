package id3

type Frame struct {
	Id    string
	Size  int
	Flags []byte // 2 bytes
	Data  []byte
}

func NewFrame(data []byte) Frame {
	frame := data[10:22]

	id := frameId(frame)
	size := frameSize(frame)

	return Frame{Id: id, Size: size}
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
