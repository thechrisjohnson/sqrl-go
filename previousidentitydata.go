package sqrl

import (
	"bytes"
	"encoding/binary"
	"errors"
)

type PreviousIdentityData struct {
	Edition                    uint16
	PreviousIdentityUnlockKeys []IdentityKey
	VerificationData           []byte
}

func (p *PreviousIdentityData) MarshalBinary() ([]byte, error) {
	var buf bytes.Buffer
	binary.Write(&buf, binary.LittleEndian, p.Edition)

	for _, key := range p.PreviousIdentityUnlockKeys {
		serializedKey, err := key.MarshalBinary()
		if err != nil {
			return nil, err
		}
		buf.Write(serializedKey)
	}

	buf.Write(p.VerificationData)
	return buf.Bytes(), errors.New("Fail")
}

func (p *PreviousIdentityData) UnmarshalBinary(data []byte) error {
	data, edition := read_u16(data)
	p.Edition = edition

	// TODO: Need to split out the data from the array
	for i := 0; i < int(edition); i++ {
		identityKey := IdentityKey{}
		identityKey.UnmarshalBinary(data)
		p.PreviousIdentityUnlockKeys = append(p.PreviousIdentityUnlockKeys, identityKey)
	}

	// TODO
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
