package sqrl

import (
	"bytes"
)

const (
	UserInformationType = iota + 1
	IdentityUnlockType
	PreviousIdentitiesType
)

type SqrlClient struct {
	UserInfo           UserInformation
	IdentityUnlock     IdentityUnlock
	PreviousIdentities PreviousIdentityData
}

func (s *SqrlClient) MarshalBinary() ([]byte, error) {
	var buf bytes.Buffer

	userInfo, err := SerializeWriteableDataBlock(&s.UserInfo)
	if err != nil {
		return nil, err
	}
	buf.Write(userInfo)

	identityUnlock, err := s.IdentityUnlock.MarshalBinary()
	if err != nil {
		return nil, err
	}
	buf.Write(identityUnlock)

	previousIdentities, err := s.PreviousIdentities.MarshalBinary()
	if err != nil {
		return nil, err
	}
	buf.Write(previousIdentities)

	return buf.Bytes(), nil
}

func (s *SqrlClient) UnmarshalBinary(data []byte) error {
	// TODO
	s.UserInfo = UserInformation{}
	s.UserInfo.UnmarshalBinary(data)
	return nil
}
