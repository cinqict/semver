package main

import "testing"

func checkInt(t *testing.T, expect, result int) {
	if result != expect {
		t.Errorf("Expected '%d' but got '%d'", expect, result)
	}
}

func TestNewSemVer(t *testing.T) {
	semver := NewSemVer(1, 2, 3)
	expectMajor := 1
	expectMinor := 2
	expectPatch := 3

	checkInt(t, expectMajor, semver.Major)
	checkInt(t, expectMinor, semver.Minor)
	checkInt(t, expectPatch, semver.Patch)

}

func TestParseSemVer(t *testing.T) {
	versionStrGood := "1.2.3"
	versionStrBad1 := "1.2.3.4"
	versionStrBad2 := "1.2"

	// Good version
	semver, _ := ParseSemVer(versionStrGood)
	expectMajor := 1
	expectMinor := 2
	expectPatch := 3

	checkInt(t, expectMajor, semver.Major)
	checkInt(t, expectMinor, semver.Minor)
	checkInt(t, expectPatch, semver.Patch)

	// Long version
	_, err1 := ParseSemVer(versionStrBad1)
	if err1 == nil {
		t.Errorf("Expect this version %s to raise an error", versionStrBad1)
	}

	// Short version
	_, err2 := ParseSemVer(versionStrBad2)
	if err2 == nil {
		t.Errorf("Expect this version %s to raise an error", versionStrBad2)
	}

	// ToDo:
	// Check for bad versions like: "v1.2.3", "1.v2.3", "1v.2.3" etc
}

func TestParseSemVerSlice(t *testing.T) {
	versions := []string{"1.2.3", "2.3.4"}
	semvers, _ := ParseSemVerSlice(versions)
	expectMajor := 2
	expectMinor := 3
	expectPatch := 4

	checkInt(t, expectMajor, semvers[1].Major)
	checkInt(t, expectMinor, semvers[1].Minor)
	checkInt(t, expectPatch, semvers[1].Patch)

	// ToDo:
	// Add error cases
}

func TestHighestSemVer(t *testing.T) {
	versions := []string{"1.2.3", "2.3.4"}
	semvers, _ := ParseSemVerSlice(versions)
	expectMajor := 2
	expectMinor := 3
	expectPatch := 4

	highest, _ := GetHighestSemVerFromSlice(semvers)

	checkInt(t, expectMajor, highest.Major)
	checkInt(t, expectMinor, highest.Minor)
	checkInt(t, expectPatch, highest.Patch)

	// Sorting an empty array
	emptySemver := []SemVer{}
	highest, error := GetHighestSemVerFromSlice(emptySemver)
	if error == nil {
		t.Errorf("Expect an empty array to raise an error")
	}
}
