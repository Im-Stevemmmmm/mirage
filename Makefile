.PHONY: clean

clean:
	rm -rf dist

run:
	go run main.go

build:
	GOOS=linux GOARCH=amd64 go build -o dist/mirage.linux
	GOOS=darwin GOARCH=amd64 go build -o dist/mirage.darwin
	GOOS=windows GOARCH=amd64 go build -o dist/mirage.exe
