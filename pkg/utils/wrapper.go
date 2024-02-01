package utils

import (
	"bytes"
	"errors"
	"fmt"
	"log"
	"os/exec"
	p "path"
)

type Wrapper struct {
	BinaryPath string
	Args       []string
	BinaryName string
}

func (w *Wrapper) ExecuteCommandWithOutput(args ...string) string {
	out, err := exec.Command(w.BinaryPath, args...).Output()
	fmt.Println(string(out))
	if err != nil {
		log.Fatal(err)
	}
	return string(out)
}

func (w *Wrapper) ExecuteCommand(args ...string) string {
	cmd := exec.Command(w.BinaryPath, args...)
	// Stdout Pipe for piping the output to variable instead to stdout
	stdout, err := cmd.StdoutPipe()
	// fmt.Println(string(out))
	if err != nil {
		log.Fatal(err)
	}
	if err := cmd.Start(); err != nil {
		log.Fatal(err)
	}
	buf := new(bytes.Buffer)
	buf.ReadFrom(stdout)
	if err := cmd.Wait(); err != nil {
		log.Fatal(err)
	}
	return buf.String()
}

func (w *Wrapper) SearchBinary(binary string) string {
	path, err := exec.LookPath(binary)
	if errors.Is(err, exec.ErrDot) {
		err = nil
	}
	if err != nil {
		log.Fatalf("'%s' binary not found in $PATH\n", binary)
	}
	log.Printf("'%s' binary available at '%s'\n", binary, path)
	w.BinaryName = p.Base(path)
	w.BinaryPath = path
	return path
}
