package main

import (
	"fmt"

	"github.com/go-git/go-git/v5/plumbing/object"
	"github.com/go-git/go-git/v5"
	u "github.com/sunshine69/golang-tools/utils"
)

func GitLog(args []string) {
	r, _ := GetWorkTree()
	logOpt := git.LogOptions{}
	cIter, err := r.Log(&logOpt)
	u.CheckErr(err, "Log")
	err = cIter.ForEach(func(c *object.Commit) error {
		fmt.Println(c)

		return nil
	})
	u.CheckErr(err, "ForEach")
}
