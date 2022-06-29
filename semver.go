package main

import (
	"fmt"
	"log"
	"os/exec"
)

// func Status() string {
// 	gitStatus := cmd.NewCmd("git", "status")
// 	return gitStatus
// }

// func main() {
// 	gitStatus, err := exec.Command("git").Output()
// 	if err != nil {
// 		fmt.Println(err)
// 	}
// 	fmt.Printf(gitStatus)
// }

func main() {
	gitStatus, err := exec.Command("git", "status").Output()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf(string(gitStatus))
}
