# Semver

## Development

### Run tests

```shell
# Unit tests
go test

# Linting
golangci-lint run # Golang lint tool: https://github.com/golangci/golangci-lint 

```

```shell
go mod tidy
go run semver.go
```

### Build

```shell
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


