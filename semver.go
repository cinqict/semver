package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

type SemVer struct {
	Major int
	Minor int
	Patch int
}

type BranchedSemVer struct {
	Version SemVer
	Branch  string
	Commit  int
}

func NewSemVer(major, minor, patch int) SemVer {
	return SemVer{
		Major: major,
		Minor: minor,
		Patch: patch,
	}
}

func (semver *SemVer) CompareTo(other SemVer) bool {
	if semver.Major != other.Major {
		return semver.Major > other.Major
	}

	if semver.Minor != other.Minor {
		return semver.Minor > other.Minor
	}

	return semver.Patch > other.Patch
}

func (semver *SemVer) IncreaseMajor() {
	semver.Major += 1
	semver.Minor = 0
	semver.Patch = 0
}

func (semver *SemVer) IncreaseMinor() {
	semver.Minor += 1
	semver.Patch = 0
}

func (semver *SemVer) IncreasePatch() {
	semver.Patch += 1
}

func NewBranchedSemVer(major int, minor int, patch int, branch string, commit int) BranchedSemVer {
	return BranchedSemVer{
		Version: SemVer{
			Major: major,
			Minor: minor,
			Patch: patch,
		},
		Branch: branch,
		Commit: commit,
	}
}

func (semver SemVer) String() string {
	return fmt.Sprintf("%d.%d.%d", semver.Major, semver.Minor, semver.Patch)
}

func ParseSemVer(versionStr string) (SemVer, error) {
	versionParts := strings.Split(versionStr, ".")

	// Check length
	if len(versionParts) != 3 {
		return SemVer{}, fmt.Errorf("Invalid sematic version format: %s", versionStr)
	}

	major, err := strconv.Atoi(versionParts[0])
	if err != nil {
		return SemVer{}, fmt.Errorf("Failed to parse major verions: %w", err)
	}

	minor, err := strconv.Atoi(versionParts[1])
	if err != nil {
		return SemVer{}, fmt.Errorf("Failed to parse minor verions: %w", err)
	}

	patch, err := strconv.Atoi(versionParts[2])
	if err != nil {
		return SemVer{}, fmt.Errorf("Failed to parse patch verions: %w", err)
	}

	return NewSemVer(major, minor, patch), nil
}

func ParseSemVerSlice(versions []string) ([]SemVer, error) {
	var parsedSemVers = []SemVer{}
	for _, version := range versions {
		parsedSemVer, err := ParseSemVer(version)
		if err != nil {
			return nil, err
		}
		parsedSemVers = append(parsedSemVers, parsedSemVer)
	}
	return parsedSemVers, nil
}

func GetHighestSemVerFromSlice(versions []SemVer) SemVer {
	sort.Slice(versions[:], func(i, j int) bool {
		return versions[i].CompareTo(versions[j])
	})
	return versions[0]
}
