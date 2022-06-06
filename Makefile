
ALL_PACKAGES= $(shell go list ./... | grep -v -e "repository" -e "mocks" -e "config")
ALL_PACKAGES_TESTS=$(shell go list ./...| grep -v "vendor")
	
stockbit-osx: main.go
	GOOS=darwin GOARCH=amd64 go build -ldflags '-s -w' -o $@

stockbit-linux: main.go
	GOOS=linux GOARCH=amd64 go build -ldflags '-s -w' -o $@

stockbit64.exe: main.go
	GOOS=windows GOARCH=amd64 go build -ldflags '-s -w' -o $@

stockbit32.exe: main.go
	GOOS=windows GOARCH=386 go build -ldflags '-s -w' -o $@

test:
	$(foreach pkg, $(ALL_PACKAGES),\
	go test -v -race $(pkg);)

test-anagram:
	go test -run TestAnagram -v
	
# Coverage with HTML output
cover-html:
	@echo "mode: count" > coverage-all.out

	$(foreach pkg, $(ALL_PACKAGES),\
	go test -coverprofile=coverage.out -covermode=count $(pkg);\
	tail -n +2 coverage.out >> coverage-all.out;)
	go tool cover -html=coverage-all.out
	rm coverage.out coverage-all.out

run:
	go run main.go echo_server.go