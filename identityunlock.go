package sqrl

type IdentityUnlock struct {
	ScryptConfig      ScryptConfig
	IdentityUnlockKey IdentityKey
	VerificationData  []byte
}

func (i *IdentityUnlock) MarshalBinary() ([]byte, error) {
	scryptConfig, err := i.ScryptConfig.MarshalBinary()
	if err != nil {
		return nil, err
	}

	identityUnlockKey, err := i.IdentityUnlockKey.MarshalBinary()
	if err != nil {
		return nil, err
	}

	return append(append(scryptConfig, identityUnlockKey...), i.VerificationData...), nil
}

func (i *IdentityUnlock) UnmarshalBinary(data []byte) error {
	i.ScryptConfig = ScryptConfig{}
	err := i.ScryptConfig.UnmarshalBinary(data)
	if err != nil {
		return err
	}

	i.IdentityUnlockKey = IdentityKey{}
	err = i.IdentityUnlockKey.UnmarshalBinary(data)
	if err != nil {
		return err
	}

	copy(i.VerificationData, data)
	return nil
}

func (i *IdentityUnlock) Length() uint16 {
	return 73
}

func (i *IdentityUnlock) Type() uint16 {
	return IdentityUnlockType
}
