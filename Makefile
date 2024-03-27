BINARY_NAME=s1t9ledcontrol

build:
	go mod tidy
	GOOS=linux GOARCH=amd64 go build -o releases/${BINARY_NAME}_linux_amd64 .
	GOOS=windows GOARCH=amd64 go build -o releases/${BINARY_NAME}_windows_amd64.exe .

clean:
	go clean
	rm -rf releases/

