package riak

import (
	"errors"
	"github.com/golang/protobuf/proto"
	"reflect"
	"time"
)

var ErrPtrRequired = errors.New("function requires a pointer argument")

type Ctrl uint8
type CAPControls map[Ctrl]uint32

const (
	N Ctrl = iota
	R
	W
	PR
	PW
	DW
)

const (
	_ = ^uint32(0) - iota

	One
	Quorum
	All
	Default
)

func setField(v reflect.Value, name string, i uint32) {
	f := v.FieldByName(name)
	if f.IsValid() {
		f.Set(reflect.ValueOf(&i))
	}
}

func SetReqCAPControls(req proto.Message, ctrl CAPControls) {
	has := func(c Ctrl) bool { _, ok := ctrl[c]; return ok }

	if ctrl == nil {
		return
	}
	v := reflect.ValueOf(req)
	if v.Kind() != reflect.Ptr {
		panic(ErrPtrRequired)
	}
	v = v.Elem()
	switch {
	case has(N):
		setField(v, "NVal", ctrl[N])
		fallthrough
	case has(R):
		setField(v, "R", ctrl[R])
		fallthrough
	case has(W):
		setField(v, "W", ctrl[W])
		fallthrough
	case has(PR):
		setField(v, "Pr", ctrl[PR])
		fallthrough
	case has(PW):
		setField(v, "Pw", ctrl[PW])
		fallthrough
	case has(DW):
		setField(v, "Dw", ctrl[DW])
	}
}

func SetReqTimeout(req proto.Message, timeout time.Duration) {
	if timeout == 0 {
		return
	}
	v := reflect.ValueOf(req)
	if v.Kind() != reflect.Ptr {
		panic(ErrPtrRequired)
	}
	v = v.Elem()
	setField(v, "Timeout", uint32(timeout.Seconds()*1000))
}
