package main

import (
	"fmt"
	"log"
	"os/exec"
)

func Status() []byte {
	result, err := exec.Command("git", "status").Output()
	if err != nil {
		log.Fatal(err)
	}
	return result
}

func Branch() []byte {
	result, err := exec.Command("git", "--no-pager", "branch", "--show-current").Output()
	if err != nil {
		log.Fatal(err)
	}
	return result
}

func AllTags() []byte {
	result, err := exec.Command("git", "--no-pager", "tag", "--list").Output()
	if err != nil {
		log.Fatal(err)
	}
	return result
}

func LatestTag() []byte {
	result, err := exec.Command("git", "describe", "--abbrev=0", "--tags").Output()
	if err != nil {
		log.Fatal(err)
	}
	return result
}

func CommitsSinceTag(tag string) []byte {
	// result, err := exec.Command("git", "rev-list", "--count", tag + "..HEAD").Output()
	fmt.Printf("HEAD", "^"+tag)
	result, err := exec.Command("git", "rev-list", "--count", "HEAD", "^"+tag).Output()
	if err != nil {
		log.Fatal(err)
	}
	return result
}

// func Log() []byte {
// 	log, err := exec.Command("git", "--no-pager", "log", "--oneline").Output()
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	return log
// }

//func main() {
//	gitStatus := Status()
//	fmt.Printf(string(gitStatus) + "\n")
//
//	branch := Branch()
//	fmt.Printf(string(branch) + "\n")
//
//	tags := AllTags()
//	fmt.Printf(string(tags) + "\n")
//
//	latestTag := LatestTag()
//	fmt.Printf(string(latestTag) + "\n")
//
//	commitsSinceTag := CommitsSinceTag(string(latestTag))
//	fmt.Printf(string(commitsSinceTag) + "\n")
//
//}
