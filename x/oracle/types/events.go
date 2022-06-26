package types

const (
	AttributeValueModule = ModuleName

	EventTypeJoin                   = "join"
	AttributeKeyID                  = "id"
	AttributeKeyOperatorAddress     = "operator_address"
	AttributeKeyEnclaveReportBase64 = "enclave_report_base64"
	AttributeKeyEncPubKeyBase64     = "enc_pub_key_base64"

	EventTypeVoteForJoin = "vote_for_join"
	AttributeKeyOption   = "option"
	AttributeKeyVoter    = "voter"

	EventTypeJoinResult = "join_result"
	AttributeKeyStatus  = "status"
	AttributeKeyValue   = "value"
)
