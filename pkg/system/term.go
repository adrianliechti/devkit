package system

import (
	"errors"
	"os"

	"golang.org/x/term"
)

func IsTerminal(d any) bool {
	if fd, ok := descriptor(d); ok {
		return term.IsTerminal(fd)
	}

	return false
}

func MakeRawTerminal(d any) (func(), error) {
	if fd, ok := descriptor(d); ok {
		state, err := term.MakeRaw(fd)

		if err != nil {
			return nil, err
		}

		return func() {
			term.Restore(fd, state)
		}, nil
	}

	return nil, errors.New("not a tty")
}

func TerminalSize(d any) (width, height int, err error) {
	if fd, ok := descriptor(d); ok {
		return term.GetSize(fd)
	}

	return 0, 0, errors.New("not a tty")
}

func descriptor(d any) (int, bool) {
	if fd, ok := d.(int); ok {
		return fd, true
	}

	if file, ok := d.(*os.File); ok {
		return int(file.Fd()), true
	}

	return 0, false
}
