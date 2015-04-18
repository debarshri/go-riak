package riak

// Performs a Riak Get request.
func (c *Conn) Get(req *RpbGetReq) (resp *RpbGetResp, err error) {
	resp = new(RpbGetResp)
	err = c.do(MsgRpbGetReq, MsgRpbGetResp, req, resp)
	return
}

// Performs a Riak Put request.
func (c *Conn) Put(req *RpbPutReq) (resp *RpbPutResp, err error) {
	resp = new(RpbPutResp)
	err = c.do(MsgRpbPutReq, MsgRpbPutResp, req, resp)
	return
}

// Performs a Riak Del request.
func (c *Conn) Del(req *RpbDelReq) error {
	return c.do(MsgRpbDelReq, MsgRpbDelResp, req, nil)
}
