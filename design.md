
# Semver CLI

## Functional description

- Get last semver/tag (default 0.0.0)
- Determine new semver
  - check branch name to determine new semver
    - if `prefix=major/*`: major+1, minor=0, patch=0
    - if `prefix=minor/*`: minor+1, patch=0
    - else patch+1
- on main
  - if PR merge, use new semver from PR branch
  - else patch+1 (i.e. direct commit to master)
- on non main branch
  - tag = new semver + branchname + number of commits since last tag

- Merge branch in main:
  Use prefix to determine increase
- Merge branch in branch:
  - patch into patch -> should work (just patch increase)
  - patch into minor -> should work (just minor increase)
  - more general: "smaller or equal" into "bigger or equal" -> should work (just "bigger" increase)
  - unusual: bigger into smaller -> this will result into "smaller" increase. Do we want that?  
- Merge main in branch:
  This should work. There is possibly a new lastTag/baseTag, prefix/branch stays the same, commit count may change.

### Examples

- direct commit to main without tags -> tag = 0.0.1
- third commit on a branch `fix/foo` with last tag 1.0.2 -> tag = 1.0.3-fix-foo-3
- merge this branch into master -> tag = 1.0.3
- merge this branch into master, but last tag already changed to 1.0.4 -> tag 1.0.5
- second commit on branch `major/bar` with last tag 3.4.8 -> tag = 4.0.0-major-bar-2
- merge this branch into master -> tag = 4.0.0

## Features

- Determine SemVer
- Determine Tag
- Set tag

### Functions

- GetLatestTag(repo) string
- StringToSemver(string) Semver
  - VerifySemver(string) bool
- GetBranchName(repo) string
- GetSemverIncrease(branchname string) string # one of major, minor, patch

### functions sketch

- getCurrentBranch(repo) string
- getDefaultBranch(repo) string
- getCurrentCommit(repo) commit
- getParrentCommitInMain(?) commit
- semverIncrement(branchname) semver
- getLatestSemverTag(commit) semver
- countCommitsSince(commit) int
- newTag(tag_current, tag_increment, currentBranch, commitCount) string
- setTag(tag_new)

## Structs

- Semver = (major int, minor int, patch int) # int or string?
  - add 4e element: (suffix string)?




