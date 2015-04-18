package riak

// Performs a Riak Get Bucket request.
func (c *Conn) GetBucket(req *RpbGetBucketReq) (resp *RpbGetBucketResp, err error) {
	resp = new(RpbGetBucketResp)
	err = c.do(MsgRpbGetBucketReq, MsgRpbGetBucketResp, req, resp)
	return
}

// Performs a Riak Set Bucket request.
func (c *Conn) SetBucket(req *RpbSetBucketReq) error {
	return c.do(MsgRpbSetBucketReq, MsgRpbSetBucketResp, req, nil)
}

// Performs a Riak List Buckets request.
// The protobufs say that it will return multiple responses but it in fact does not.
func (c *Conn) ListBuckets(req *RpbListBucketsReq) (resp *RpbListBucketsResp, err error) {
	resp = new(RpbListBucketsResp)
	err = c.do(MsgRpbListBucketsReq, MsgRpbListBucketsResp, req, resp)
	return
}

// Performs a Riak List Keys request.
// Returns multiple list keys responses.
func (c *Conn) ListKeys(req *RpbListKeysReq) ([]*RpbListKeysResp, error) {
	var resps []*RpbListKeysResp

	if err := c.request(MsgRpbListKeysReq, req); err != nil {
		return nil, err
	}
	for {
		resp := new(RpbListKeysResp)
		if err := c.response(MsgRpbListKeysResp, resp); err != nil {
			return nil, err
		}
		resps = append(resps, resp)

		if resp.GetDone() {
			break
		}
	}
	return resps, nil
}

// Performs a Riak Reset Bucket request.
func (c *Conn) ResetBucket(req *RpbResetBucketReq) error {
	return c.do(MsgRpbResetBucketReq, MsgRpbResetBucketResp, req, nil)
}

// Performs a Riak Get Bucket Type request.
func (c *Conn) GetBucketType(req *RpbGetBucketTypeReq) (resp *RpbGetBucketResp, err error) {
	resp = new(RpbGetBucketResp)
	err = c.do(MsgRpbGetBucketTypeReq, MsgRpbGetBucketResp, req, resp)
	return
}

// Performs a Riak Set Bucket Type request.
func (c *Conn) SetBucketType(req *RpbSetBucketTypeReq) error {
	return c.do(MsgRpbSetBucketTypeReq, MsgRpbSetBucketResp, req, nil)
}
