package sqrl

import "errors"

type ScryptConfig struct {
	RandomSalt      []byte
	LogNFactor      uint8
	IterationFactor uint32
}

func (s *ScryptConfig) MarshalBinary() ([]byte, error) {
	// TODO
	return nil, errors.New("Fail")
}

func (s *ScryptConfig) UnmarshalBinary([]byte) error {
	// TODO
	return nil
}
