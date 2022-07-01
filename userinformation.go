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

func (u *UserInformation) MarshalBinary() ([]byte, error) {
	return nil, errors.New("Fail")
}

func (u *UserInformation) UnmarshalBinary(data []byte) error {
	return nil
}

func (u *UserInformation) Length() uint16 {
	return 125
}

func (u *UserInformation) Type() uint16 {
	return UserInformationType
}
