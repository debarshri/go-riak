package riak

// Performs a Riak Ping request.
func (c *Conn) Ping() error {
	return c.do(MsgRpbPingReq, MsgRpbPingResp, nil, nil)
}

// Performs a Riak Server info request.
func (c *Conn) ServerInfo() (resp *RpbGetServerInfoResp, err error) {
	resp = new(RpbGetServerInfoResp)
	err = c.do(MsgRpbGetServerInfoReq, MsgRpbGetServerInfoResp, nil, resp)
	return
}

// Performs a Riak Auth request.
func (c *Conn) Authenticate(req *RpbAuthReq) error {
	return c.do(MsgRpbAuthReq, MsgRpbAuthResp, req, nil)
}

// Performs a Riak bucket-key preflist request.
func (c *Conn) BucketKeyPreflist(req *RpbGetBucketKeyPreflistReq) (resp *RpbGetBucketKeyPreflistResp, err error) {
	resp = new(RpbGetBucketKeyPreflistResp)
	err = c.do(MsgRpbGetBucketKeyPreflistReq, MsgRpbGetBucketKeyPreflistResp, req, resp)
	return
}
