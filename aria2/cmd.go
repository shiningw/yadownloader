package aria2

import (
	"errors"
	"io"
	"os/exec"
)

type Aria2Cmd struct {
	RunOptions
}

func NewAria2Cmd(opts RunOptions) *Aria2Cmd {
	a := &Aria2Cmd{opts}
	//a.Options = opts.Options
	return a
}
func (a *Aria2Cmd) Run() error {
	var bin string
	var err error
	if bin, err = exec.LookPath("aria2c"); err != nil {
		return errors.New("aria2c not found")
	}
	cmd := exec.Command(bin, a.Options...)
	stderr, err := cmd.StderrPipe()
	if err != nil {
		return err
	}
	if err := cmd.Start(); err != nil {
		return err
	}

	b, _ := io.ReadAll(stderr)
	if len(b) > 0 {
		cmd.Process.Release()
		return errors.New(string(b))
	}

	if err := cmd.Wait(); err != nil {
		return err
	}
	return nil
}

func StartAria2(opts RunOptions) error {
	a := NewAria2Cmd(opts)
	return a.Run()
}

func StopAria2(a *Aria2) error {
	return a.Shutdown()
}

func Aria2IsRunning(a *Aria2) error {

	return a.IsRunning()
}
