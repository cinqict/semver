package main

import (
	"fmt"
	"strconv"
	"strings"
)

type SemVer struct {
	Major int
	Minor int
	Patch int
}

func NewSemVer(major, minor, patch int) SemVer {
	return SemVer{
		Major: major,
		Minor: minor,
		Patch: patch,
	}
}

func (version SemVer) String() string {
	return fmt.Sprintf("%d.%d.%d", version.Major, version.Minor, version.Patch)
}

func ParseSemVer(versionStr string) (SemVer, error) {
	versionParts := strings.Split(versionStr, ".")

	// Check lenght
	if len(versionParts) != 3 {
		return SemVer{}, fmt.Errorf("Invalid sematic version format: %s", versionStr)
	}

	major, err := strconv.Atoi(versionParts[0])
	if err != nil {
		return SemVer{}, fmt.Errorf("Failed to parse major verions: %w", err)
	}

	minor, err := strconv.Atoi(versionParts[1])
	if err != nil {
		return SemVer{}, fmt.Errorf("Failed to parse major verions: %w", err)
	}

	patch, err := strconv.Atoi(versionParts[2])
	if err != nil {
		return SemVer{}, fmt.Errorf("Failed to parse major verions: %w", err)
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
