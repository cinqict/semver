package main

import (
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
)

func GetRepo(path string) *git.Repository {
	repo, err := git.PlainOpen(path)
	CheckIfError(err)
	return repo
}

func GetAllTags(repo *git.Repository) []string {
	tagIter, err := repo.Tags()
	CheckIfError(err)

	var tagRefs []string
	err = tagIter.ForEach(func(tagRef *plumbing.Reference) error {
		tagRefs = append(tagRefs, tagRef.Name().Short())
		return nil
	})
	CheckIfError(err)
	return tagRefs
}
