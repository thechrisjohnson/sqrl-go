package sqrl

import "errors"

type IdentityKey struct {
	Key []byte
}

func (s *IdentityKey) MarshalBinary() ([]byte, error) {
	return nil, errors.New("Fail")
}
