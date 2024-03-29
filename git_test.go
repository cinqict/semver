package main

import (
	"reflect"
	"sort"
	"testing"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
	"github.com/go-git/go-git/v5/storage/memory"
)

// Create a temporary in-memory Git repository
func InitTestRepo() *git.Repository {
	repo, err := git.Init(memory.NewStorage(), nil)
	CheckIfError(err)

	return repo
}

// Create some tags for testing
func AddTestTags(repo *git.Repository, tags []string) *git.Repository {
	for _, tag := range tags {
		_, err := repo.CreateTag(tag, plumbing.NewHash("some-commit-hash"), nil)
		CheckIfError(err)
	}
	return repo
}

func TestGetAllTags(t *testing.T) {
	repo := InitTestRepo()
	tagNames := []string{"v1.0.0", "v1.1.0", "v2.0.0"}
	repoWithTags := AddTestTags(repo, tagNames)

	// Call the function you want to test
	tagResult := GetAllTags(repoWithTags)
	tagExpected := tagNames

	sort.Strings(tagExpected)
	sort.Strings(tagResult)

	// Assert the result
	if !reflect.DeepEqual(tagExpected, tagResult) {
		t.Errorf("Tags are not equal. Expected: %v, Actual: %v", tagExpected, tagResult)
	}
}
