package riak

import (
	"bytes"
)

// Performs a Riak Index (2i) request.
// The protobufs say that it will return multiple responses but it in fact does not.
func (c *Conn) Index(req *RpbIndexReq) (resp *RpbIndexResp, err error) {
	resp = new(RpbIndexResp)
	err = c.do(MsgRpbIndexReq, MsgRpbIndexResp, req, resp)
	return
}

// Performs a Riak Search Query request.
func (c *Conn) SearchQuery(req *RpbSearchQueryReq) (resp *RpbSearchQueryResp, err error) {
	resp = new(RpbSearchQueryResp)
	err = c.do(MsgRpbSearchQueryReq, MsgRbpSearchQueryResp, req, resp)
	return
}

// Performs a Riak Map Reduce request.
// Returns multiple map-reduce responses.
func (c *Conn) MapRed(req *RpbMapRedReq) ([]*RpbMapRedResp, error) {
	var resps []*RpbMapRedResp

	if err := c.request(MsgRpbMapRedReq, req); err != nil {
		return nil, err
	}
	for {
		resp := new(RpbMapRedResp)
		if err := c.response(MsgRpbMapRedResp, resp); err != nil {
			return nil, err
		}
		resps = append(resps, resp)

		if resp.GetDone() {
			break
		}
	}
	return resps, nil
}

// GetMany is a convenience method that uses map-reduce with
// Riak built-in erlang functions to get many documents at once.
//
// Warning: If you specify an empty list of keys, it will do a full
// map-reduce on the bucket (usually not a good idea).
func (c *Conn) GetMany(bucket string, keys []string) ([][]byte, error) {
	var res [][]byte

	req := &RpbMapRedReq{
		Request:     setUnionQuery(bucket, keys),
		ContentType: []byte("application/json"),
	}
	resps, err := c.MapRed(req)
	if err != nil {
		return nil, err
	}

	for _, resp := range resps {
		if resp.Response == nil {
			continue
		}
		res = append(res, resp.GetResponse())
	}
	return res, nil
}

func setUnionQuery(bucket string, keys []string) []byte {
	b := bytes.NewBuffer([]byte(`{"inputs":`))
	if len(keys) == 0 {
		b.WriteRune('"')
		b.WriteString(bucket)
		b.WriteRune('"')
	} else {
		b.WriteRune('[')
		for i, key := range keys {
			if i > 0 {
				b.WriteRune(',')
			}
			b.Write([]byte(`["` + bucket + `","` + key + `"]`))
		}
		b.WriteRune(']')
	}
	b.Write([]byte(`,"query":[` +
		`{"map":{"language":"erlang","module":"riak_kv_mapreduce","function":"map_object_value"}},` +
		`{"reduce":{"language":"erlang","module":"riak_kv_mapreduce","function":"reduce_set_union"}}]}`))
	return b.Bytes()
}
