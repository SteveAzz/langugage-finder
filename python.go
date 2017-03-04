package main

import (
	"bytes"
	"log"
	"os/exec"
)

// Python represent information about php
type Python struct {
	Name    string
	Version string
	Path    string
}

// GetName php
func (lang Python) GetName() string {
	if lang.Name != "" {
		return lang.Name
	}

	return "python"
}

// GetVersion gets the current version of php active in Path
// This will take only in consideration the the php cli environment
func (lang Python) GetVersion() string {
	path := lang.GetPath()

	if path == "" {
		return ""
	}
	// get php version
	var out bytes.Buffer

	cmd := exec.Command(path, "-r", "echo PHP_VERSION;")
	cmd.Stdout = &out

	err := cmd.Run()
	if err != nil {
		log.Println(err)
	}

	return out.String()
}

// GetPath gets the path of the php version
func (lang Python) GetPath() string {
	// get the path first
	path, installed := exec.LookPath("php")

	// If it is not installed
	if installed != nil {
		return ""
	}

	return path
}
