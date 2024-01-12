package main

import "golang.org/x/mod/semver"

func GetHighestSemVer(tags []string) string {
	semver.Sort(tags)
	return tags[0]
}
