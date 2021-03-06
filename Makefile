build:
	GOOS=darwin GOARCH=386 go build -o bin/jenkinsctl-darwin_i386 main.go
	GOOS=linux GOOARCH=amd64 go build -o bin/jenkinsctl-x86_64 main.go
depend:
	dep ensure -vendor-only
fmt:
	gofmt -w cmd output pkg
	goimports -w cmd output pkg
run:
	go run main.go
test:
	go test -v ./...
cover:
	go test -cover -v ./...
test-docker:
	docker build -f build/test.Dockerfile .
