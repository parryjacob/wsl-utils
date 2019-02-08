package main

import (
	"os"
	"os/exec"
	"strings"
	"syscall"

	log "github.com/Sirupsen/logrus"
)

var wslExePath = "C:\\Windows\\System32\\wsl.exe"

func wslpath(p string) string {
	//p = strings.Replace(p, "\\", "\\\\", -1)

	cmd := quietBat("wslpath", p)
	bs, err := cmd.Output()
	if err != nil {
		log.Error(string(bs))
		panic(err)
	}
	wslp := strings.TrimSpace(string(bs))

	return wslp
}

func quietBat(p string, args ...string) *exec.Cmd {
	cmdPath := os.Getenv("WINDIR") + "\\System32\\cmd.exe"
	nargs := []string{"/c", p}
	nargs = append(nargs, args...)
	cmd := exec.Command(cmdPath, nargs...)

	// This will only compile on Windows
	cmd.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}

	return cmd
}

func main() {
	args := []string{}

	if len(os.Args) < 2 {
		log.Fatal("You must provide at least 1 argument!")
	}

	for i, v := range os.Args {
		if i == 0 || i == 1 {
			continue
		}

		if len(v) > 3 && v[1:3] == ":\\" {
			// this is probably a Windows path
			args = append(args, wslpath(v))
			continue
		}

		args = append(args, v)
	}

	cmd := quietBat(os.Args[1], args...)

	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	cmd.Start()
}
