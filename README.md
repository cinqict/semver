# Semver

## Development

### Run tests

```shell
# Unit tests
go test

# Linting
golangci-lint run # Golang lint tool: https://github.com/golangci/golangci-lint
# or run linting in docker
docker run --rm -v $(pwd):/app -w /app docker/golangci-lint:1.55.1-go1.21.6 golangci-lint run # Workaround: docker/golangci-lint:1.55.1-go1.21.6 image fixes golangci/golangci-lint:v1.55 old version issue

```

### Dev Run

```shell
go run .
```


### Update dependecies

```shell
go mod tidy
```

### Build

```shell
go build
go install

```

