PKG_NAME=linksaver

mod:
	cd src && \
	rm go.mod | true && \
	go mod init $(PKG_NAME) && \
	go mod tidy

run:
	cd src && \
	go run .

go-build:
	cd src && \
	go build .
