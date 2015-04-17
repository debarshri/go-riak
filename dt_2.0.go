package riak

// Performs a Riak CRDT Fetch request.
func (c *Conn) DtFetch(req *DtFetchReq) (resp *DtFetchResp, err error) {
	resp = new(DtFetchResp)
	err = c.do(MsgDtFetchReq, MsgDtFetchResp, req, resp)
	return
}

// Performs a Riak CRDT Update request.
func (c *Conn) DtUpdate(req *DtUpdateReq) (resp *DtUpdateResp, err error) {
	resp = new(DtUpdateResp)
	err = c.do(MsgDtUpdateReq, MsgDtUpdateResp, req, resp)
	return
}
