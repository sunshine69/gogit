package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
	"syscall"

	"path/filepath"

	"golang.org/x/term"
)

var (
	action string // clone, commit, log, push pull
)

func main() {
	var sshkeyfile string
	flag.StringVar(&sshkeyfile, "ssh-key", os.Getenv("HOME")+"/.ssh/id_rsa", "ssh key file path")
	flag.Parse()

	args := flag.Args()
	fmt.Println(args)
	action = args[0]
	fmt.Println(action)
	switch action {
	case "clone":
		giturl := args[1]
		fmt.Print("Enter Password: ")
		bytePassword, err := term.ReadPassword(int(syscall.Stdin))
		if err != nil {
			panic(err)
		}
		password := string(bytePassword)
		directory := ""
		if len(args) == 3 {
			directory = os.Args[2]
		} else {
			directory = filepath.Dir(giturl)
			directory = strings.TrimSuffix(directory, ".git")
		}
		Clone(giturl, directory, sshkeyfile, password)
	}

}
