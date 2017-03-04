package main

import (
	"bytes"
	"log"
	"os/exec"
	"strings"
)

// Node represent information about php
type Node struct {
	Name    string
	Version string
	Path    string
}

// GetName php
func (lang Node) GetName() string {
	if lang.Name != "" {
		return lang.Name
	}

	return "node"
}

// GetVersion gets the current version of php active in Path
// This will take only in consideration the the php cli environment
func (lang Node) GetVersion() string {
	path := lang.GetPath()

	if path == "" {
		return ""
	}
	// get php version
	var out bytes.Buffer

	cmd := exec.Command(path, "-v")
	cmd.Stdout = &out

	err := cmd.Run()
	if err != nil {
		log.Println(err)
	}

	return strings.Trim(out.String(), "v")
}

// GetPath gets the path of the php version
func (lang Node) GetPath() string {
	// get the path first
	path, installed := exec.LookPath("node")

	// If it is not installed
	if installed != nil {
		return ""
	}

	return path
}
