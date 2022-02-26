# precommit check
precommit: goimports fmt errcheck staticcheck gosec

goimports:
	find . -print | grep --regex '.*\.go' | xargs goimports -w -local "github.com/hirac1220/go-file/utils"

fmt:
	go fmt ./...

vet:
	go vet ./...

errcheck:
	errcheck

staticcheck:
	staticcheck ./...

golangci:
	golangci-lint run ./...

gosec:
	gosec ./...


# unit test
test_all: test_file

test_file:
	go test -v -coverpkg=. ./...

test_cover_file:
	go test -v -coverprofile=cover_file.out  -coverpkg=./...
	go tool cover -html=cover_file.out -o cover_file.html
