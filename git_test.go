package main

import (
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
	"github.com/go-git/go-git/v5/storage/memory"
	"reflect"
	"testing"
)

// Create a temporary in-memory Git repository
func InitRepo() *git.Repository {
	repo, err := git.Init(memory.NewStorage(), nil)
	CheckIfError(err)

	return repo
}

// Create some tags for testing
func AddTags(repo *git.Repository, tags []string) *git.Repository {
	for _, tag := range tags {
		//tagRef := plumbing.NewTagReferenceName(tagName)
		_, err := repo.CreateTag(tag, plumbing.NewHash("some-commit-hash"), nil)
		CheckIfError(err)
	}
	return repo
}

func TestGetAllTags(t *testing.T) {
	repo := InitRepo()
	tagNames := []string{"v1.0.0", "v1.1.0", "v2.0.0"}
	repoWithTags := AddTags(repo, tagNames)

	// Call the function you want to test
	tagResult := GetAllTags(repoWithTags)
	tagExpected := tagNames

	// Assert the result
	if !reflect.DeepEqual(tagExpected, tagResult) {
		t.Errorf("Tags are not equal. Expected: %v, Actual: %v", tagExpected, tagResult)
	}
}
