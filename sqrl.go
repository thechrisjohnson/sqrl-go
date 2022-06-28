package sqrl

import "errors"

type SqrlClient struct {
	UserInfo       UserInformation
	IdentityUnlock IdentityUnlock
}

func (s *SqrlClient) MarshalBinary() ([]byte, error) {
	s.UserInfo.MarshalBinary()
	s.IdentityUnlock.MarshalBinary()
	return nil, errors.New("Fail")
}
