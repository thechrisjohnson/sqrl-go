package sqrl

import "errors"

type IdentityUnlock struct {
	ScryptConfig      ScryptConfig
	IdentityUnlockKey IdentityKey
	VerificationData  []byte
}

func (i *IdentityUnlock) MarshalBinary() ([]byte, error) {
	return nil, errors.New("Fail")
}
