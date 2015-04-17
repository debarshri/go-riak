package riak

// Performs a Riak Auth request.
func (c *Conn) Authenticate(req *RpbAuthReq) error {
	return c.do(MsgRpbAuthReq, MsgRpbAuthResp, req, nil)
}
