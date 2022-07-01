package sqrl

import (
	"bytes"
)

type IdentityKey struct {
	Key []byte
}

func (s *IdentityKey) MarshalBinary() ([]byte, error) {
	var buf bytes.Buffer
	_, err := buf.Write(s.Key)
	if err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

func (s *IdentityKey) UnmarshalBinary(data []byte) error {
	copy(s.Key, data)
	return nil
}
