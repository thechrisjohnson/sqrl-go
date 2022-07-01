package sqrl

import "errors"

type PreviousIdentityData struct {
	Edition                    uint16
	PreviousIdentityUnlockKeys []IdentityKey
	VerificationData           []byte
}

func (p *PreviousIdentityData) MarshalBinary() ([]byte, error) {
	return nil, errors.New("Fail")
}

func (p *PreviousIdentityData) UnmarshalBinary(data []byte) error {
	return nil
}

func (p *PreviousIdentityData) Length() uint16 {
	if p.Edition > 0 {
		return 22 + (p.Edition * 32)
	} else {
		return 0
	}
}

func (p *PreviousIdentityData) Type() uint16 {
	return PreviousIdentitiesType
}
