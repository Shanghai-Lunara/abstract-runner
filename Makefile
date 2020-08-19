.PHONY: abs

abs:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o abstract-runner cmd/main.go

abs2:
	rm -rf abstract-runner
	go build -o abstract-runner cmd/main.go



