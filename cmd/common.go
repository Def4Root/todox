package cmd

import (
	"fmt"
	"os"
	"path/filepath"
)

const (
	todoListFile = ".todox"
)

func GetListFile() (listFile string) {
	home := os.Getenv("HOME")
	if home == "" {
		os.Getenv("USERPROFILE")
	}
	listFile = filepath.Join(home, todoListFile)
	return
}

func FailOnError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}
}