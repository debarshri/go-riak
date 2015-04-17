package riak

const (
	MsgRpbResetBucketReq         = 29
	MsgRpbResetBucketResp        = 30
	MsgRpbGetBucketTypeReq       = 31 // Resp code: MsgRpbGetBucketResp
	MsgRpbSetBucketTypeReq       = 32 // Resp code: MsgRpbSetBucketResp
	MsgRpbYokozunaIndexGetReq    = 54
	MsgRpbYokozunaIndexGetResp   = 55
	MsgRpbYokozunaIndexPutReq    = 56 // Resp code: MsgRpbPutResp
	MsgRpbYokozunaIndexDeleteReq = 57 // Resp code: MsgRpbDelResp
	MsgRpbYokozunaSchemaGetReq   = 58
	MsgRpbYokozunaSchemaGetResp  = 59
	MsgRpbYokozunaSchemaPutReq   = 60 // Resp code: MsgRpbPutResp
	MsgDtFetchReq                = 80
	MsgDtFetchResp               = 81
	MsgDtUpdateReq               = 82
	MsgDtUpdateResp              = 83
	MsgRpbAuthReq                = 253
	MsgRpbAuthResp               = 254
	MsgRpbStartTls               = 255 // Riak security (TLS required)
)
