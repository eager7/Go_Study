package daemon

import (
	"fmt"
	"github.com/sevlyar/go-daemon"
)

func Initialize() (*daemon.Context, error) {
	cnTxt := &daemon.Context{
		PidFileName: "pid",
		PidFilePerm: 0644,
		LogFileName: "log",
		LogFilePerm: 0644,
		WorkDir:     "./",
		Chroot:      "",
		Env:         nil,
		Args:        []string{"[go-daemon sample]"},
		Credential:  nil,
		Umask:       027,
	}

	parent, err := cnTxt.Reborn()
	if err != nil {
		return nil, err
	}
	if parent != nil {
		fmt.Println("the parent process will exit...")
		return nil, nil
	}
	fmt.Println("run child process")
	return cnTxt, nil
}
