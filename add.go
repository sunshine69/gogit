package main

import (
	"flag"
)

func GitAdd(args []string) {
	_, w := GetWorkTree()

	addFlag := flag.NewFlagSet("clone opt", flag.ExitOnError)
	addAll := addFlag.Bool("a", false, "Add all modified files")
	globPtn := addFlag.String("g", "", "File glob pattern to add")

	addFlag.Parse(args)

	if *addAll {
		w.AddGlob("**")
		return
	}

	if *globPtn != "" {
		w.AddGlob(*globPtn)
		return
	}

	fileList := addFlag.Args()
	if len(fileList) == 0 {
		_, fileList = Status([]string{})
	}

	for _, _f := range fileList {
		w.Add(_f)
	}

}
