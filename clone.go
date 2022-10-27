package main

import (
	"fmt"
	"os"

	git "github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing/transport/ssh"
	u "github.com/sunshine69/golang-tools/utils"
)

func Clone(url, directory, privateKeyFile, password string) {
	fmt.Printf("url: %s - directory: %s key: %s\n", url, directory, privateKeyFile)

	_, err := os.Stat(privateKeyFile)
	u.CheckErr(err, "privateKeyFile")

	// Clone the given repository to the given directory
	fmt.Printf("git clone %s\n", url)
	publicKeys, err := ssh.NewPublicKeysFromFile("git", privateKeyFile, password)
	u.CheckErr(err, "NewPublicKeysFromFile")

	r, err := git.PlainClone(directory, false, &git.CloneOptions{
		// The intended use of a GitHub personal access token is in replace of your password
		// because access tokens can easily be revoked.
		// https://help.github.com/articles/creating-a-personal-access-token-for-the-command-line/
		Auth:     publicKeys,
		URL:      url,
		Progress: os.Stdout,
	})
	u.CheckErr(err, "PlainClone")

	// ... retrieving the branch being pointed by HEAD
	ref, err := r.Head()
	u.CheckErr(err, "Head")
	// ... retrieving the commit object
	commit, err := r.CommitObject(ref.Hash())

	u.CheckErr(err, "CommitObject")

	fmt.Println(commit)
}
