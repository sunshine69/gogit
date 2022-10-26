package main

import (
	"flag"
	"fmt"
	"strings"

	u "github.com/sunshine69/golang-tools/utils"
)

type StatusCode int8

const (
	CLEAN = iota
	// At least one file modifed and in staging area. Can commit
	MODIFIED_STAGING
	// All files modifed but non of them in staging. Do not commit
	MODIFIED_NONSTAGING
)

// Return statusCode and list o modified files but not staged
func Status(args []string) (StatusCode, []string) {
	_, w := GetWorkTree()

	statusFlag := flag.NewFlagSet("status opt", flag.ExitOnError)
	optShort := statusFlag.String("s", "", "Short status")
	branchStatus := statusFlag.String("b", "", "Branch status")
	fmt.Printf("%s %s\n", *optShort, *branchStatus)
	statusFlag.Parse(args)

	status, err := w.Status()
	u.CheckErr(err, "Status")

	if status.String() != "" {
		statusMap := ParseStatus(status.String())
		listModifiedNonStagedFiles := []string{}
		listModifiedStagedFiles := []string{}
		for file := range statusMap {
			fs := status.File(file)
			if fs.Staging == 'M' {
				listModifiedStagedFiles = append(listModifiedStagedFiles, file)
			} else {
				listModifiedNonStagedFiles = append(listModifiedNonStagedFiles, file)
			}
		}
		stagedFilesCount, nonStagedFilesCount := len(listModifiedStagedFiles), len(listModifiedNonStagedFiles)
		if nonStagedFilesCount > 0 {
			fmt.Println("Modified files not staged")
			for _, f := range listModifiedNonStagedFiles {
				fmt.Printf("%s %s\n", statusMap[f], f)
			}
		}
		if stagedFilesCount > 0 {
			fmt.Println("Modified files staged")
			for _, f := range listModifiedStagedFiles {
				fmt.Printf("%s %s\n", statusMap[f], f)
			}
		}
		if nonStagedFilesCount == 0 {
			return MODIFIED_STAGING, []string{}
		}
		if stagedFilesCount == 0 && nonStagedFilesCount > 0 {
			return MODIFIED_NONSTAGING, listModifiedNonStagedFiles
		}
		return MODIFIED_STAGING, listModifiedNonStagedFiles
	}
	// fmt.Printf("is clean: %v\n", status.IsClean())
	return CLEAN, []string{}
}

func ParseStatus(statusString string) map[string]string {
	sLines := strings.Split(statusString, "\n")
	o := map[string]string{}
	for _, l := range sLines {
		l = strings.TrimSpace(l)
		// fmt.Println(l)
		status_file_lst := strings.Fields(l)
		if len(status_file_lst) == 2 {
			o[status_file_lst[1]] = status_file_lst[0]
		}
	}
	return o
}
