package main

import (
	"github.com/go-git/go-git/v5"
	"golang.org/x/mod/semver"
)

func GetHighestSemVer(repo *git.Repository) string {
	tags := GetAllTags(repo)
	semver.Sort(tags)
	return tags[0]
}
