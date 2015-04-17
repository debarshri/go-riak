package riak

import (
	"github.com/3XX0/pooly"
	"net"
	"regexp"
	"time"
)

var hostport = regexp.MustCompile("[?.*]?:.*")

type Driver struct {
	connTimeout  time.Duration
	writeTimeout time.Duration
	readTimeout  time.Duration
	testInterval time.Duration
}

func Client(c *pooly.Conn) *Conn {
	return c.Interface().(*Conn)
}

func NewDriver() *Driver {
	return new(Driver)
}

func (d *Driver) SetConnTimeout(t time.Duration) {
	d.connTimeout = t
}

func (d *Driver) SetReadTimeout(t time.Duration) {
	d.readTimeout = t
}

func (d *Driver) SetWriteTimeout(t time.Duration) {
	d.writeTimeout = t
}

func (d *Driver) SetTestInterval(t time.Duration) {
	d.testInterval = t
}

func (d *Driver) Dial(address string) (*pooly.Conn, error) {
	var c net.Conn
	var err error

	if !hostport.MatchString(address) {
		address += ":8087" // default riak pbc port
	}

	if d.connTimeout > 0 {
		c, err = net.DialTimeout("tcp", address, d.connTimeout)
	} else {
		c, err = net.Dial("tcp", address)
	}
	if err != nil {
		return nil, err
	}

	conn := &Conn{
		conn:         c.(*net.TCPConn),
		writeTimeout: d.writeTimeout,
		readTimeout:  d.readTimeout,
	}
	return pooly.NewConn(conn), nil
}

func (d *Driver) Close(conn *pooly.Conn) {
	c := Client(conn)
	if c != nil {
		c.conn.Close()
	}
}

func (d *Driver) TestOnBorrow(conn *pooly.Conn) error {
	c := Client(conn)
	if !c.lastChecked.IsZero() {
		if t := time.Now().Sub(c.lastChecked); t < d.testInterval {
			return nil
		}
	}
	c.lastChecked = time.Now()

	return c.Ping()
}

func (d *Driver) Temporary(err error) bool {
	if e, ok := err.(net.Error); ok {
		return e.Temporary()
	}
	return false
}
