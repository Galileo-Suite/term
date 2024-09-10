//go:build !windows && !aix
// +build !windows,!aix

package term

import "golang.org/x/sys/unix"

func ioctlGetWinsize(fd int) (*unix.Winsize, error) {
	return unix.IoctlGetWinsize(fd, unix.TIOCGWINSZ)
}

func ioctlSetWinsize(fd int, value *unix.Winsize) error {
	return unix.IoctlSetWinsize(fd, unix.TIOCSWINSZ, value)
}
