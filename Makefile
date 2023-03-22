lint: 
		go mod tidy
		go fmt ./... 
		go vet ./...

test:
	go clean -testcache
	go test -v ./...

test-coverage:
	go test -coverprofile cover.out ./... && go tool cover -html=cover.out
