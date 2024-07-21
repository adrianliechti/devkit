package moby

import (
	"errors"
	"os"

	"golang.org/x/term"
)

func IsTerminal(fd any) bool {
	if file, ok := fd.(*os.File); ok {
		return term.IsTerminal(int(file.Fd()))
	}

	return false
}

func MakeRawTerminal(fd any) (func(), error) {
	if file, ok := fd.(*os.File); ok {
		state, err := term.MakeRaw(int(file.Fd()))

		if err != nil {
			return nil, err
		}

		return func() {
			term.Restore(int(file.Fd()), state)
		}, nil
	}

	return nil, errors.New("not a tty")
}

func TerminalSize(fd any) (width, height int, err error) {
	if file, ok := fd.(*os.File); ok {
		return term.GetSize(int(file.Fd()))
	}

	return 0, 0, errors.New("not a tty")
}
