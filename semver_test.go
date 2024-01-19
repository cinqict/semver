package main

import (
	"reflect"
	"testing"
)

func TestGetHighestSemVer(t *testing.T) {
	repo := InitRepo()
	tagNames := []string{"v1.0.9", "v1.3.0", "v2.0.0"}
	repoWithTags := AddTags(repo, tagNames)

	tagResult := GetHighestSemVer(repoWithTags)
	tagExpected := "v2.0.0"

	// Assert the result
	//assert.ElementsMatch(t, tagExpected, tagResult)
	if !reflect.DeepEqual(tagExpected, tagResult) {
		t.Errorf("Not the highest tag. Expected: %v, Actual: %v", tagExpected, tagResult)
	}
}
