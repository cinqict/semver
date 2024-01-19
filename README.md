# Semver

## Development

```shell
go mod tidy
go run semver.go

go test
golangci-lint run # Golang lint tool: https://github.com/golangci/golangci-lint 
go build
go install

```

## ToDo

- [ ] log
- [x] current branch
- [x] tags
- [ ] count commits
- [x] latest tag
- [ ] commits since tag
- [ ] unit tests
- [x] use git module
  - [x] show git log
- [ ] make it a cli with `github.com/urfave/cli/v2`
- [x] golinting


