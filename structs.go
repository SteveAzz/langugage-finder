package main

// Language information about the language
type Language interface {
	GetName() string
	GetVersion() string
	GetPath() string
}

// Lang struct representing information about language for json payload
type Lang struct {
	Name    string
	Version string
	Path    string
}

// Languages a slice of all the users installed languages
type Languages struct {
	Languages []Language
}
