package moby

import (
	"io"
	"os"

	"golang.org/x/term"
)

func IsTTY(stdin io.Reader) bool {
	if file, ok := stdin.(*os.File); ok {
		return term.IsTerminal(int(file.Fd()))
	}

	return false
}

func TTYSize() (width, height int, err error) {
	return term.GetSize(0)
}
