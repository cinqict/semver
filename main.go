package main

import (
	"fmt"
)

func main() {
	// Get repo
	repo := GetRepo(".")

	// Get all tags
	allTags := GetAllTags(repo)
	fmt.Println(allTags)

	//// SemVer to string
	//semver1 := NewSemVer(1, 2, 0)
	//version1 := semver1.String()
	//fmt.Printf("SemVer: %s\n", version1)
	//
	//// String to SemVer
	//version2 := "1.2.3"
	//semver2, err := ParseSemVer(version2)
	//CheckIfError(err)
	//fmt.Printf("SemVer major: %d, minor: %d, patch %d\n", semver2.Major, semver2.Minor, semver2.Patch)

	// Array / slice of versions (allTags) to semver
	allTagsSemVer, err := ParseSemVerSlice(allTags)
	CheckIfError(err)
	for _, semver := range allTagsSemVer {
		fmt.Printf("SemVer major: %d, minor: %d, patch %d\n", semver.Major, semver.Minor, semver.Patch)
	}

	highest, _ := GetHighestSemVerFromSlice(allTagsSemVer)
	fmt.Println(highest)

	highest.IncreaseMinor()

	fmt.Println(highest)
}
