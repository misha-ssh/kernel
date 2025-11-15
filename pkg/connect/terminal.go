package connect

import (
	"os"

	"github.com/misha-ssh/kernel/internal/logger"
	"golang.org/x/crypto/ssh"
	"golang.org/x/term"
)

// todo put file ssh pkg
const (
	Timeout = 0

	EnableMod = 1
	ICRNLMod  = 1
	INLCRMod  = 1
	ISIGMod   = 1

	ISPEED = 115200
	OSPEED = 115200

	TypeTerm       = "xterm-256color"
	HeightTerminal = 80
	WidthTerminal  = 40
)

func createTerminalSession(session *ssh.Session) error {
	fd := int(os.Stdin.Fd())
	width, height, err := term.GetSize(fd)
	if err != nil {
		width = WidthTerminal
		height = HeightTerminal
	}

	modes := ssh.TerminalModes{
		ssh.ECHO:          EnableMod,
		ssh.ICRNL:         ICRNLMod,
		ssh.INLCR:         INLCRMod,
		ssh.ISIG:          ISIGMod,
		ssh.TTY_OP_ISPEED: ISPEED,
		ssh.TTY_OP_OSPEED: OSPEED,
	}

	if err = session.RequestPty(TypeTerm, height, width, modes); err != nil {
		logger.Error(err.Error())
		return err
	}

	session.Stdin = os.Stdin
	session.Stdout = os.Stdout
	session.Stderr = os.Stderr

	return nil
}
