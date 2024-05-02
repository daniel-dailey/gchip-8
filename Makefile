build:
	go mod tidy
	go build -o dist/gochip8 cmd/main.go

clean:
	rm -rf dist