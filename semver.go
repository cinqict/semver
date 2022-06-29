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

// func Log() []byte {
// 	log, err := exec.Command("git", "--no-pager", "log", "--oneline").Output()
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	return log
// }

// - [ ]  log
// - [x] current branch
// - [ ]  tags
// - [ ]  count commits
// - [ ]  latest tag
// - [ ]  commits since tag

func main() {
	gitStatus := Status()
	fmt.Printf(string(gitStatus))

	branch := Branch()
	fmt.Printf(string(branch))

}
