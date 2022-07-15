//go:build !darwin && !linux
package peercred

import (
	"syscall"
)

func getRawPeercred(c syscall.RawConn) (*Peercred, error) {
	return nil, ErrNotImplemented
}
