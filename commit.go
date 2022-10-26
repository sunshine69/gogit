package main

import (
	"flag"
	"fmt"

	"time"

	git "github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing/object"
	u "github.com/sunshine69/golang-tools/utils"
)

func Commit(gitUser, gitUserEmail string, args []string) {
	commitFlag := flag.NewFlagSet("clone opt", flag.ExitOnError)
	commitMsg := commitFlag.String("m", "", "Commit msg")
	commitFlag.Parse(args)

	_, w := GetWorkTree()

	statusCode, _ := Status([]string{})
	// fmt.Printf("status %v\n", statusCode)

	switch statusCode {
	case CLEAN:
		fmt.Println("Clean, do nothing")
		return
	case MODIFIED_NONSTAGING:
		fmt.Println("All files are not staged. Use git add files to stage them first")
		return
	case MODIFIED_STAGING:
		_, err := w.Commit(*commitMsg, &git.CommitOptions{
			Author: &object.Signature{
				Name:  gitUser,
				Email: gitUserEmail,
				When:  time.Now(),
			},
		})
		u.CheckErr(err, "Commit")
	}

}
