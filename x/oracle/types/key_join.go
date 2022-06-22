package types

import "encoding/binary"

var _ binary.ByteOrder

const (
	// JoinKeyPrefix is the prefix to retrieve all Join
	JoinKeyPrefix = "Join/value/"
)

// JoinKey returns the store key to retrieve a Join from the index fields
func JoinKey(
	joinID uint64,
) []byte {
	var key []byte

	idBytes := make([]byte, 8)
	binary.LittleEndian.PutUint64(idBytes, joinID)
	key = append(key, idBytes...)
	key = append(key, []byte("/")...)

	return key
}
