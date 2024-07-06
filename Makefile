.PHONY: all clean test 


.PHONY: help
help:
	@echo 'Usage:'
	@sed -n 's/^##//p' ${MAKEFILE_LIST} | column -t -s ':' |  sed -e 's/^/ /'

.PHONY: confirm
confirm:
	@echo -n 'Are you sure? [y/N] ' && read ans && [ $${ans:-N} = y ]

.PHONY: no-dirty
no-dirty:
	git diff --exit-code


# =================
# Quality Checks
# =================
## tidy: format code and tidy modfile
.PHONY: tidy
tidy:
	@go fmt ./...
	@go mod tidy



## check: checks for errors/vulnerabilities in the code
.PHONY: check
check:
	@go run honnef.co/go/tools/cmd/staticcheck@latest -checks=all,-ST1000,-U1000 ./...
	@go run golang.org/x/vuln/cmd/govulncheck@latest ./...
	@go run github.com/kisielk/errcheck@latest ./...

## audit: run quality control checks
.PHONY: audit
audit: check
	@go mod verify
	@go vet ./...
	@go test -race -buildvcs -vet=off ./...

# ============================
# Development
# ============================

## test: run all tests
.PHONY: test
test:
	go test -v -race -buildvcs ./...

.PHONY: test/cover
test/cover:
	go test -race -buildvcs -cover -coverprofile=./tmp/coverage.out ./...
	go tool cover -html=./tmp/coverage.out


