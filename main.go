package main

import (
	"fmt"
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing/object"
	"k8s.io/klog/v2"
	"log"
)

func main() {
	var path string
	path = "."
	repo, err := git.PlainOpen(path)
	CheckIfError(err)

	// Gets the HEAD history from HEAD, just like this command:
	klog.Infoln("git log")

	// ... retrieves the branch pointed by HEAD
	// Get the HEAD reference
	ref, err := repo.Head()
	if err != nil {
		log.Fatal(err)
	}

	// Retrieve the commit history starting from the HEAD reference
	commits, err := repo.Log(&git.LogOptions{
		From: ref.Hash(),
	})
	if err != nil {
		log.Fatal(err)
	}

	// Iterate over the commit history and print commit information
	err = commits.ForEach(func(commit *object.Commit) error {
		fmt.Printf("Commit Hash: %s\n", commit.Hash)
		fmt.Printf("Author: %s <%s>\n", commit.Author.Name, commit.Author.Email)
		fmt.Printf("Date: %s\n", commit.Author.When)
		fmt.Printf("Message: %s\n", commit.Message)
		fmt.Println("")

		return nil
	})
	if err != nil {
		log.Fatal(err)
	}
}
