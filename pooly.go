package riak

import (
	"github.com/3XX0/pooly"
	"log"
	"net"
	"regexp"
	"time"
)

var hostport = regexp.MustCompile("[?.*]?:.*")

type Driver struct {
	logger       *log.Logger
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

func (d *Driver) SetLogger(l *log.Logger) {
	d.logger = l
}

func (d *Driver) logf(format string, v ...interface{}) {
	if d.logger != nil {
		d.logger.Printf(format+"\n", v...)
	}
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
		d.logf("dial error: %v", err)
		return nil, err
	}

	conn := &Conn{
		conn:         c.(*net.TCPConn),
		writeTimeout: d.writeTimeout,
		readTimeout:  d.readTimeout,
	}
	d.logf("connection opened (%s)", address)
	return pooly.NewConn(conn), nil
}

func (d *Driver) Close(conn *pooly.Conn) {
	c := Client(conn)
	if c != nil {
		c.conn.Close()
	}
	d.logf("connection closed")
}

func (d *Driver) TestOnBorrow(conn *pooly.Conn) error {
	c := Client(conn)
	if !c.lastChecked.IsZero() {
		if t := time.Now().Sub(c.lastChecked); t < d.testInterval {
			return nil
		}
	}
	c.lastChecked = time.Now()

	err := c.Ping()
	if err != nil {
		d.logf("borrowing failed: %v", err)
	}
	return err
}

func (d *Driver) Temporary(err error) bool {
	var tmp bool

	if e, ok := err.(net.Error); ok {
		tmp = e.Temporary()
	}
	if !tmp {
		d.logf("fatal connection error: %v", err)
	}
	return tmp
}
