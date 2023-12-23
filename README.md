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

```

## ToDo

- [ ] log
- [x] current branch
- [x] tags
- [ ] count commits
- [x] latest tag
- [ ] commits since tag
- [ ] unit tests
- [ ] use git module
  - [x] show git log
- [ ] make it a cli with `github.com/urfave/cli/v2`
- [x] golinting

## Hackaton Semver cli

Prep:
 
- [ ] Cli basics
- [ ] Linting
- [ ] Unit test: test repo
- [ ] Unittest: example
- [ ] Create go files/modules
- [ ] Create structs
- [ ] eigen idee van functies ()

Prep collega's

- [ ] Go function
- [ ] Go struct

Agenda
- Intro: idee, show basics
- Design functions (allemaal)
  - input, output, go file, pseudo algorithm
- Tweetallen: kies een functie
  - check
    - unit tests
    - linting

## Ref

- https://go.dev/blog/using-go-modules
- https://pkg.go.dev/os/exec#Cmd.Output