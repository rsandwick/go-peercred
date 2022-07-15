package peercred // "rs3.io/go/peercred"

import (
	"fmt"
	"net"
	"runtime"
)

var (
	ErrNotImplemented      = fmt.Errorf("not implemented on " + runtime.GOOS)
	ErrUnsupportedConnType = fmt.Errorf("unsupported connection type")
)

type Peercred struct{ Pid, Uid, Gid int }

func Get(c net.Conn) (*Peercred, error) {
	x, ok := c.(*net.UnixConn)
	if !ok {
		return nil, ErrUnsupportedConnType
	}
	raw, err := x.SyscallConn()
	if err != nil {
		return nil, fmt.Errorf("SyscallConn: %w", err)
	}
	return getRawPeercred(raw)
}
