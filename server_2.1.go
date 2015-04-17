package riak

// Performs a Riak bucket-key preflist request.
func (c *Conn) BucketKeyPreflist(req *RpbGetBucketKeyPreflistReq) (resp *RpbGetBucketKeyPreflistResp, err error) {
	resp = new(RpbGetBucketKeyPreflistResp)
	err = c.do(MsgRpbGetBucketKeyPreflistReq, MsgRpbGetBucketKeyPreflistResp, req, resp)
	return
}
