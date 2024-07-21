package moby

import (
	"errors"
	"io"
	"os"

	"golang.org/x/term"
)

func IsTerminal(fd io.Reader) bool {
	if file, ok := fd.(*os.File); ok {
		return term.IsTerminal(int(file.Fd()))
	}

	return false
}

func MakeRawTerminal(fd any) (*term.State, error) {
	if file, ok := fd.(*os.File); ok {
		return term.MakeRaw(int(file.Fd()))
	}

	return nil, errors.New("not a tty")
}

func RestoreTerminal(fd any, state *term.State) error {
	if file, ok := fd.(*os.File); ok {
		return term.Restore(int(file.Fd()), state)
	}

	return errors.New("not a tty")
}

func TerminalSize(fd any) (width, height int, err error) {
	if file, ok := fd.(*os.File); ok {
		return term.GetSize(int(file.Fd()))
	}

	return 0, 0, errors.New("not a tty")
}
