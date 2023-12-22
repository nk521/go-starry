package util

import (
	"os"
	"os/exec"
	"runtime"
)

func OpenEditor(path string) error {
	opener := "vi" // default to vi

	if s := os.Getenv("EDITOR"); s != "" {
		opener = s
	} else {
		switch runtime.GOOS {
		case "windows":
			opener = "notepad.exe"
		}
	}

	cmd := exec.Command(opener, path)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err := cmd.Run()
	return err
}
