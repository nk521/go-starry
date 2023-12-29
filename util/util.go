package util

import (
	"bufio"
	"log"
	"net/http"
	"os"
	"os/exec"
	"runtime"
	"strings"
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

func ParseHeadersFromString(headerStr string) *http.Header {
	headerStr = strings.Replace(headerStr, "HTTP/2", "HTTP/2.0", 1)
	headerStr = strings.ReplaceAll(headerStr, "\n", "\r\n")
	reader := bufio.NewReader(strings.NewReader(headerStr + "\r\n"))
	logReq, err := http.ReadRequest(reader)
	if err != nil {
		log.Fatal(err)
	}
	return &logReq.Header
}
