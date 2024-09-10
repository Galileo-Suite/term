//go:build aix
// +build aix

package term

import "golang.org/x/sys/unix"

func ioctlGetWinsize(fd int) (*unix.Winsize, error) {
	return unix.IoctlGetWinsize(fd, unix.TIOCGWINSZ)
}

func ioctlSetWinsize(fd int, value *unix.Winsize) error {
	return unix.IoctlSetWinsize(fd, unix.TIOCSWINSZ&0xffffffff, value)
}
