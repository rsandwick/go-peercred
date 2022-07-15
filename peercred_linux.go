package peercred

import (
	"fmt"
	"syscall"

	"golang.org/x/sys/unix"
)

func getRawPeercred(c syscall.RawConn) (*Peercred, error) {
	var cred *unix.Ucred
	var err error
	controlErr := raw.Control(func(fd uintptr) {
		cred, err = unix.GetsockoptXucred(
			int(fd),
			unix.SOL_SOCKET,
			unix.SO_PEERCRED,
		)
	})
	if controlErr != nil {
		return nil, fmt.Errorf("raw.Control: %w", controlErr)
	}
	if err != nil {
		return nil, fmt.Errorf("unix.GetsockoptUcred: %w", err)
	}
	return &Peercred{
		Pid: int(cred.Pid),
		Uid: int(cred.Uid),
		Gid: int(cred.Gid),
	}, nil
}
