package sqrl

import (
	"bytes"
	"errors"
)

const (
	UserInformationType = iota + 1
	IdentityUnlockType
	PreviousIdentitiesType
)
const FileHeader = "sqrldata"

type SqrlClient struct {
	UserInfo           UserInformation
	IdentityUnlock     IdentityUnlock
	PreviousIdentities PreviousIdentityData
}

func (s *SqrlClient) MarshalBinary() ([]byte, error) {
	var buf bytes.Buffer

	// Start with the file header
	buf.WriteString(FileHeader)

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
	// Confirm the first eight characters of the binary data match the file header
	data, header := read_string(data, 8)
	if header != FileHeader {
		return errors.New("Invalid file header")
	}

	// While we have more data to read
	for len(data) > 0 {
		data, blockLength := read_u16(data)
		if len(data) < int(blockLength) {
			return errors.New("Invalid binary data: Mismatched block length and data size")
		}

		// Split out this block's data
		blockData := data[:blockLength]
		data = data[blockLength:]

		blockData, blockType := read_u16(blockData)
		hasUserInfo, hasIdentityUnlock := false, false
		switch blockType {
		case UserInformationType:
			if hasUserInfo {
				return errors.New("Invalid binary data: Multiple user informtaion blocks found")
			}

			s.UserInfo = UserInformation{}
			s.UserInfo.UnmarshalBinary(blockData)
			hasUserInfo = true
			break
		case IdentityUnlockType:
			if hasIdentityUnlock {
				return errors.New("Invalid binary data: Multiple identity unlock blocks found")
			}

			s.IdentityUnlock = IdentityUnlock{}
			s.IdentityUnlock.UnmarshalBinary(blockData)
			hasIdentityUnlock = true
			break
		case PreviousIdentitiesType:
			s.PreviousIdentities = PreviousIdentityData{}
			s.PreviousIdentities.UnmarshalBinary(blockData)
			break
		default:
			return errors.New("Invalid binary data: Invalid block type")
		}

		if !hasUserInfo {
			return errors.New("Invalid binary data: No user information block found")
		}
		if !hasIdentityUnlock {
			return errors.New("Invalid binary data: No identity unlock block found")
		}
	}
	// Read the type of the data we have
	// Read the size of the data we have
	s.UserInfo = UserInformation{}
	s.UserInfo.UnmarshalBinary(data)
	return nil
}
