package sqrl

import (
	"bytes"
	"encoding/binary"
)

type WriteableDataBlock interface {
	Length() uint16
	Type() uint16
	MarshalBinary() ([]byte, error)
}

func SerializeWriteableDataBlock(w WriteableDataBlock) ([]byte, error) {
	var buf bytes.Buffer
	length := w.Length()
	if length > 0 {
		// First we write the length
		err := binary.Write(&buf, binary.LittleEndian, length)
		if err != nil {
			return nil, err
		}

		// Then write the type of the object
		err = binary.Write(&buf, binary.LittleEndian, w.Type())
		if err != nil {
			return nil, err
		}

		// Then we write the actual data
		serialized, err := w.MarshalBinary()
		if err != nil {
			return nil, err
		}

		buf.Write(serialized)
	}

	return buf.Bytes(), nil
}
