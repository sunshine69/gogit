package main

import (
	"github.com/go-git/go-git/v5"
)
// TODO add more options and test
func GitPush(args []string) {
	r, _ := GetWorkTree()
	pushOpt := git.PushOptions{}
	r.Push(&pushOpt)
}
