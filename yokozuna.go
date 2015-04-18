package riak

// Perform a Riak Yokozuna Index Get request.
func (c *Conn) YokozunaIndexGet(req *RpbYokozunaIndexGetReq) (resp *RpbYokozunaIndexGetResp, err error) {
	resp = new(RpbYokozunaIndexGetResp)
	err = c.do(MsgRpbYokozunaIndexGetReq, MsgRpbYokozunaIndexGetResp, req, resp)
	return
}

// Perform a Riak Yokozuna Index Put request.
func (c *Conn) YokozunaIndexPut(req *RpbYokozunaIndexPutReq) error {
	return c.do(MsgRpbYokozunaIndexPutReq, MsgRpbPutResp, req, nil)
}

// Perform a Riak Yokozuna Index Delete request.
func (c *Conn) YokozunaIndexDelete(req *RpbYokozunaIndexDeleteReq) error {
	return c.do(MsgRpbYokozunaIndexDeleteReq, MsgRpbDelResp, req, nil)
}

// Perform a Riak Yokozuna Index Get request.
func (c *Conn) YokozunaSchemaGet(req *RpbYokozunaSchemaGetReq) (resp *RpbYokozunaSchemaGetResp, err error) {
	resp = new(RpbYokozunaSchemaGetResp)
	err = c.do(MsgRpbYokozunaSchemaGetReq, MsgRpbYokozunaSchemaGetResp, req, resp)
	return
}

// Perform a Riak Yokozuna Schema Put request.
func (c *Conn) YokozunaSchemaPut(req *RpbYokozunaSchemaPutReq) error {
	return c.do(MsgRpbYokozunaSchemaPutReq, MsgRpbPutResp, req, nil)
}
