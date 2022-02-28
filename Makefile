
install:
	cd ./cmd/runshell; go install

test:
	go test -v ./internal/command
	go test -v ./internal/file
