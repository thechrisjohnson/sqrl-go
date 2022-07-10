package sqrl

import "encoding/binary"

func read_u16(data []byte) ([]byte, uint16) {
	number := binary.LittleEndian.Uint16(data[:2])
	data = data[2:]
	return data, number
}

func read_string(data []byte, length uint) ([]byte, string) {
	text := string(data[:length])
	data = data[length:]
	return data, text
}
