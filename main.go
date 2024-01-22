package main

import (
	"fmt"
)

func main() {
	repo := GetRepo(".")
	allTags := GetAllTags(repo)
	fmt.Printf("%v\n", allTags)

	v := NewSemVer(1, 2, 0)
	vv := v.String()
	fmt.Printf("SemVer: %v\n", vv)

	//// Gets the HEAD history from HEAD, just like this command:
	//klog.Infoln("git log")
	//
	//// ... retrieves the branch pointed by HEAD
	//// Get the HEAD reference
	//ref, err := repo.Head()
	//if err != nil {
	//	log.Fatal(err)
	//}
	//
	//// Retrieve the commit history starting from the HEAD reference
	//commits, err := repo.Log(&git.LogOptions{
	//	From: ref.Hash(),
	//})
	//if err != nil {
	//	log.Fatal(err)
	//}
	//
	//// Iterate over the commit history and print commit information
	//err = commits.ForEach(func(commit *object.Commit) error {
	//	fmt.Printf("Commit Hash: %s\n", commit.Hash)
	//	fmt.Printf("Author: %s <%s>\n", commit.Author.Name, commit.Author.Email)
	//	fmt.Printf("Date: %s\n", commit.Author.When)
	//	fmt.Printf("Message: %s\n", commit.Message)
	//	fmt.Println("")
	//
	//	return nil
	//})
	//if err != nil {
	//	log.Fatal(err)
	//}
}
