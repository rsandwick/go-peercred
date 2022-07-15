package peercred

import (
	"fmt"
	"syscall"

	"golang.org/x/sys/unix"
)

func getRawPeercred(c syscall.RawConn) (*Peercred, error) {
	var (
		cred *unix.Xucred
		err  error
		pid  int
	)
	controlErr := c.Control(func(fd uintptr) {
		cred, err = unix.GetsockoptXucred(
			int(fd),
			unix.SOL_LOCAL,
			unix.LOCAL_PEERCRED,
		)
		if err != nil {
			err = fmt.Errorf("unix.GetsockoptXucred: %w", err)
			return
		}
		pid, err = unix.GetsockoptInt(
			int(fd),
			unix.SOL_LOCAL,
			unix.LOCAL_PEERPID,
		)
		if err != nil {
			err = fmt.Errorf("unix.GetsockoptInt: %w", err)
		}
	})
	if controlErr != nil {
		return nil, fmt.Errorf("raw.Control: %w", controlErr)
	}
	if err != nil {
		return nil, err
	}
	return &Peercred{
		Pid: pid,
		Uid: int(cred.Uid),
	}, nil
}
