package types

import "encoding/binary"

var _ binary.ByteOrder

var (
	NextJoinIDKey             = []byte{0x00}
	JoinKeyPrefix             = []byte{0x01}
	PendingJoinQueueKeyPrefix = []byte{0x02}
	VoteForJoinKeyPrefix      = []byte{0x03}
	OracleKeyPrefix           = []byte{0x04}
	OraclePubKeyKey           = []byte{0x05}
)

// JoinKey returns the store key to retrieve a Join from the index fields
func JoinKey(joinID uint64) []byte {
	return GetJoinIDBytes(joinID)
}

func PendingJoinQueueKey(joinID uint64) []byte {
	return GetJoinIDBytes(joinID)
}

// VoteForJoinKey returns the store key to retrieve a VoteForJoin from the index fields
func VoteForJoinKey(joinID uint64, voter string) []byte {
	return append(VoteForJoinsKey(joinID), []byte(voter)...)
}

// VoteForJoinsKey gets the 1st part of the votes key based on the joinID
func VoteForJoinsKey(joinID uint64) []byte {
	return GetJoinIDBytes(joinID)
}

func GetJoinIDBytes(joinID uint64) []byte {
	idBytes := make([]byte, 8)
	binary.LittleEndian.PutUint64(idBytes, joinID)
	return idBytes
}

func GetJoinIDFromBytes(bz []byte) uint64 {
	return binary.LittleEndian.Uint64(bz)
}

func SplitPendingJoinQueueKey(key []byte) uint64 {
	return GetJoinIDFromBytes(key)
}

func OracleKey(operatorAddress string) []byte {
	return []byte(operatorAddress)
}
