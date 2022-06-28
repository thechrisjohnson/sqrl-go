package sqrl

import "errors"

type ScryptConfig struct {
}

func (s *ScryptConfig) MarshalBinary() ([]byte, error) {
	return nil, errors.New("Fail")
}
