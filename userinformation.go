package sqrl

import "errors"

type UserInformation struct {
	AesGcmIv           []byte
	ScryptConfig       ScryptConfig
	OptionFlags        uint16
	HintLength         uint8
	PasswordVerifySecs uint8
	IdleTimeoutMins    uint16
	IdentityMasterKey  IdentityKey
	IdentityLockKey    IdentityKey
	VerificationData   []byte
}

func (s *UserInformation) MarshalBinary() ([]byte, error) {
	return nil, errors.New("Fail")
}
