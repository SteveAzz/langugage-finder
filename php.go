package main

import (
	"bytes"
	"log"
	"os/exec"
)

// Php represent information about php
type Php struct {
	Name    string
	Version string
	Path    string
}

// GetName php
func (lang Php) GetName() string {
	if lang.Name != "" {
		return lang.Name
	}

	return "php"
}

// GetVersion gets the current version of php active in Path
// This will take only in consideration the the php cli environment
func (lang Php) GetVersion() string {
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
func (lang Php) GetPath() string {
	// get the path first
	path, installed := exec.LookPath("php")

	// If it is not installed
	if installed != nil {
		return ""
	}

	return path
}
