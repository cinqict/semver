package main

import (
	"fmt"
	"log"
	"os/exec"
)

func Status() []byte {
	status, err := exec.Command("git", "status").Output()
	if err != nil {
		log.Fatal(err)
	}
	return status
}

func Branch() []byte {
	branch, err := exec.Command("git", "--no-pager", "branch", "--show-current").Output()
	if err != nil {
		log.Fatal(err)
	}
	return branch
}

func AllTags() []byte {
	branch, err := exec.Command("git", "--no-pager", "tag", "--list").Output()
	if err != nil {
		log.Fatal(err)
	}
	return branch
}


// func Log() []byte {
// 	log, err := exec.Command("git", "--no-pager", "log", "--oneline").Output()
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	return log
// }

// - [ ]  log
// - [x] current branch
// - [x]  tags
// - [ ]  count commits
// - [ ]  latest tag
// - [ ]  commits since tag

func main() {
	gitStatus := Status()
	fmt.Printf(string(gitStatus) + "\n")

	branch := Branch()
	fmt.Printf(string(branch) + "\n")

	tags := AllTags()
	fmt.Printf(string(tags) + "\n")
}
