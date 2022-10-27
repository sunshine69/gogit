package main

import (
	"flag"
	"fmt"
	"net/url"
	"os"
	"path/filepath"
	"strings"

	git "github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing/transport/http"
	"github.com/go-git/go-git/v5/plumbing/transport/ssh"
	u "github.com/sunshine69/golang-tools/utils"
	gitu "github.com/whilp/git-urls"
)

func Clone(privateKeyFile string, args []string) {
	cloneFlag := flag.NewFlagSet("clone opt", flag.ExitOnError)
	gitUrl := args[0]
	var urlp *url.URL
	var err error
	directory := ""

	if len(args) > 2 && !strings.HasPrefix(args[1], "-") {
		directory = args[1]
	} else {
		urlp, err = gitu.Parse(gitUrl)
		u.CheckErr(err, "url parse")
		// fmt.Println(urlp.Path)
		directory = filepath.Base(urlp.Path)
		// fmt.Println(directory)
		directory = strings.TrimSuffix(directory, ".git")
	}
	// we do not have any cloneOpt yet so not in use
	cloneFlag.Parse(args[1:])

	// fmt.Printf("url: %s - directory: %s key: %s\n", gitUrl, directory, privateKeyFile)
	username, password := "", ""
	cloneOpt := git.CloneOptions{
		URL:      gitUrl,
		Progress: os.Stdout,
	}
	if urlp.Scheme == "https" {
		// The token case is that username can be anything and password is the token. Fall into this code
		username = urlp.User.Username()
		if username != "" {
			if _password, ok := urlp.User.Password(); ok {
				password = _password
			} else {
				password = GetPassword()
			}
			cloneOpt.Auth = &http.BasicAuth{
				Username: username,
				Password: password,
			}
		}
	} else {
		_, err := os.Stat(privateKeyFile)
		u.CheckErr(err, "privateKeyFile")
		password = GetPassword()
		publicKeys, err := ssh.NewPublicKeysFromFile("git", privateKeyFile, password)
		u.CheckErr(err, "NewPublicKeysFromFile")
		cloneOpt.Auth = publicKeys
	}

	// Clone the given repository to the given directory
	fmt.Printf("git clone %s\n", gitUrl)

	r, err := git.PlainClone(directory, false, &cloneOpt)
	u.CheckErr(err, "PlainClone")

	// ... retrieving the branch being pointed by HEAD
	ref, err := r.Head()
	u.CheckErr(err, "Head")
	// ... retrieving the commit object
	commit, err := r.CommitObject(ref.Hash())

	u.CheckErr(err, "CommitObject")

	fmt.Println(commit)
}
