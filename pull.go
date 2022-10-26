package main

import "github.com/go-git/go-git/v5"

func GitPull(args []string) {
	_,w:= GetWorkTree()
	pullOpt := git.PullOptions{}
	w.Pull(&pullOpt)
}