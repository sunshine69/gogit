package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"syscall"

	git "github.com/go-git/go-git/v5"
	u "github.com/sunshine69/golang-tools/utils"
	gitu "github.com/whilp/git-urls"
	"golang.org/x/crypto/ssh/terminal"
	"gopkg.in/ini.v1"
)

var (
	action string // clone, commit, log, push pull
)

func ParseGitConfig(gitUserConfigFile string) *ini.File {
	if gitUserConfigFile == "" {
		gitUserConfigFile = os.Getenv("HOME") + "/.gitconfig"
	}
	cfg, err := ini.Load(gitUserConfigFile)
	u.CheckErr(err, "Load")
	return cfg
}

func main() {
	var sshkeyfile, gitConfig string
	flag.StringVar(&sshkeyfile, "ssh-key", os.Getenv("HOME")+"/.ssh/id_rsa", "ssh key file path")
	// flag.StringVar(&commitMsg, "m", "Auto commit", "Commit message")
	flag.StringVar(&gitConfig, "c", os.Getenv("HOME")+"/.gitconfig", "Git user config file")
	flag.Parse()
	cfg := ParseGitConfig(gitConfig)

	args := flag.Args()
	if len(args) == 0 {
		panic("Usage: gitg <action> <option>.")
	}

	action = args[0]

	switch action {
	case "clone":
		cloneFlag := flag.NewFlagSet("clone opt", flag.ExitOnError)
		gitUrl := args[1]
		password := GetPassword()
		directory := ""
		if !strings.HasPrefix(args[2], "-") {
			directory = args[2]
		} else {
			urlp, err := gitu.Parse(gitUrl)
			u.CheckErr(err, "url parse")
			fmt.Println(urlp.Path)
			directory = filepath.Base(urlp.Path)
			fmt.Println(directory)
			directory = strings.TrimSuffix(directory, ".git")
		}
		// we do not have any cloneOpt yet so not in use
		cloneFlag.Parse(args[1:])

		Clone(gitUrl, directory, sshkeyfile, password)
	case "commit":
		gitUser := cfg.Section("user").Key("name").String()
		gitUserEmail := cfg.Section("user").Key("email").String()

		Commit(gitUser, gitUserEmail, args[1:])
	case "status":
		Status(args[1:])
	case "add":
		GitAdd(args[1:])
	case "push":
		GitPush(args[1:])
	case "pull":
		GitPull(args[1:])
	case "log":
		GitLog(args[1:])
	}

}

func GetPassword() string {
	fmt.Println("Enter Password:")
	bytePassword, err := terminal.ReadPassword(int(syscall.Stdin))
	u.CheckErr(err, "ReadPassword")
	return string(bytePassword)
}

func GetWorkTree() (*git.Repository, *git.Worktree)  {
	workDir, err := os.Getwd()
	u.CheckErr(err, "workDir")
	r, err := git.PlainOpen(workDir)
	u.CheckErr(err, "PlainOpen")
	w, err := r.Worktree()
	u.CheckErr(err, "Worktree")
	return r, w
}
