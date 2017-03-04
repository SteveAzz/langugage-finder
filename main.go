package main

import "fmt"

// Find all the installed languages the user has
func main() {
	// Php
	php := Php{Name: "php"}
	php.Version = php.GetVersion()
	php.Path = php.GetPath()

	// Ruby
	ruby := Ruby{Name: "ruby"}
	ruby.Version = ruby.GetVersion()
	ruby.Path = ruby.GetPath()

	// Python
	python := Python{Name: "python"}
	python.Version = python.GetVersion()
	python.Path = python.GetPath()

	// Node
	node := Node{Name: "node"}
	node.Version = node.GetVersion()
	node.Path = node.GetPath()

	langs := Languages{
		Languages: []Language{
			php,
			ruby,
			python,
			node,
		},
	}

	fmt.Println(langs)
}
