package protocol

import (
	"encoding/hex"
	"io"
	"math"
	"strconv"

	"github.com/kolosok86/go-steam/protocol/steamlang"
)

type JobId uint64
type TargetJobName string
type Realm uint32
type EErrorMessage string

func (j JobId) String() string {
	if j == math.MaxUint64 {
		return "(none)"
	}
	return strconv.FormatUint(uint64(j), 10)
}

type Serializer interface {
	Serialize(w io.Writer) error
}

type Deserializer interface {
	Deserialize(r io.Reader) error
}

type Serializable interface {
	Serializer
	Deserializer
}

type MessageBody interface {
	Serializable
	GetEMsg() steamlang.EMsg
}

// the default details to request in most situations
const EClientPersonaStateFlag_DefaultInfoRequest = steamlang.EClientPersonaStateFlag_PlayerName |
	steamlang.EClientPersonaStateFlag_Presence | steamlang.EClientPersonaStateFlag_SourceID |
	steamlang.EClientPersonaStateFlag_GameExtraInfo

const DefaultAvatar = "fef49e7fa7e1997310d705b2a6158ff8dc1cdfeb"

func ValidAvatar(avatar []byte) bool {
	str := hex.EncodeToString(avatar)
	return !(str == "0000000000000000000000000000000000000000" || len(str) != 40)
}
