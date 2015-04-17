package riak

import (
	"errors"
	"github.com/golang/protobuf/proto"
	"io"
	"net"
	"time"
)

var ErrInvalidResponseCode = errors.New("invalid response code")

type Conn struct {
	conn         *net.TCPConn
	lastChecked  time.Time
	writeTimeout time.Duration
	readTimeout  time.Duration
}

// Encodes a request code and proto structure into a message byte buffer.
func encode(code uint8, req proto.Message) (buf []byte, err error) {
	if req != nil {
		buf, err = proto.Marshal(req)
		if err != nil {
			return
		}
	}
	size := uint32(len(buf) + 1)
	header := []byte{byte(size >> 24), byte(size >> 16), byte(size >> 8), byte(size), code}
	buf = append(header, buf...)
	return
}

// Decodes a message byte buffer into a proto response, error code or nil.
// Resulting object depends on response type.
func decode(expect uint8, buf []byte, resp proto.Message) error {
	if len(buf) < 1 {
		return ErrInvalidResponseCode
	}
	code := uint8(buf[0])
	buf = buf[1:]

	if code == MsgRpbErrorResp {
		resp = new(RpbErrorResp)
	} else if code != expect {
		return ErrInvalidResponseCode
	}
	if resp == nil {
		return nil
	}
	err := proto.Unmarshal(buf, resp)

	e, ok := resp.(*RpbErrorResp)
	if ok && err == nil {
		err = errors.New(string(e.Errmsg))
	}
	return err
}

// Encode and write a request to the Riak server.
func (c *Conn) request(code uint8, req proto.Message) error {
	buf, err := encode(code, req)
	if err != nil {
		return err
	}
	if c.writeTimeout > 0 {
		c.conn.SetWriteDeadline(time.Now().Add(c.writeTimeout))
	}
	_, err = c.conn.Write(buf)
	return err
}

// Read and decode a response from the Riak server.
func (c *Conn) response(expect uint8, resp proto.Message) error {
	if c.readTimeout > 0 {
		c.conn.SetReadDeadline(time.Now().Add(c.readTimeout))
	}

	buf := make([]byte, 4)
	if _, err := io.ReadFull(c.conn, buf); err != nil {
		return err
	}
	size := uint32(buf[0])<<24 | uint32(buf[1])<<16 | uint32(buf[2])<<8 | uint32(buf[3])

	buf = make([]byte, size)
	if _, err := io.ReadFull(c.conn, buf); err != nil {
		return err
	}

	return decode(expect, buf, resp)
}

func (c *Conn) do(code, expect uint8, req, resp proto.Message) error {
	if err := c.request(code, req); err != nil {
		return err
	}
	return c.response(expect, resp)
}
