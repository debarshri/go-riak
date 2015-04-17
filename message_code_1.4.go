package riak

const (
	MsgRpbErrorResp         = 0
	MsgRpbPingReq           = 1
	MsgRpbPingResp          = 2
	MsgRpbGetClientIdReq    = 3 // XXX Deprecated in v1.4
	MsgRpbGetClientIdResp   = 4 // XXX Deprecated in v1.4
	MsgRpbSetClientIdReq    = 5 // XXX Deprecated in v1.4
	MsgRpbSetClientIdResp   = 6 // XXX Deprecated in v1.4
	MsgRpbGetServerInfoReq  = 7
	MsgRpbGetServerInfoResp = 8
	MsgRpbGetReq            = 9
	MsgRpbGetResp           = 10
	MsgRpbPutReq            = 11
	MsgRpbPutResp           = 12
	MsgRpbDelReq            = 13
	MsgRpbDelResp           = 14
	MsgRpbListBucketsReq    = 15
	MsgRpbListBucketsResp   = 16
	MsgRpbListKeysReq       = 17
	MsgRpbListKeysResp      = 18
	MsgRpbGetBucketReq      = 19
	MsgRpbGetBucketResp     = 20
	MsgRpbSetBucketReq      = 21
	MsgRpbSetBucketResp     = 22
	MsgRpbMapRedReq         = 23
	MsgRpbMapRedResp        = 24
	MsgRpbIndexReq          = 25
	MsgRpbIndexResp         = 26
	MsgRpbSearchQueryReq    = 27
	MsgRbpSearchQueryResp   = 28
	MsgRpbCSBucketReq       = 40 // Riak CS only
	MsgRpbCSBucketResp      = 41 // Riak CS only
	MsgRpbCounterUpdateReq  = 50 // Counters v1.4
	MsgRpbCounterUpdateResp = 51 // Counters v1.4
	MsgRpbCounterGetReq     = 52 // Counters v1.4
	MsgRpbCounterGetResp    = 53 // Counters v1.4
)
