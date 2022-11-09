PKG_NAME=linksaver
DUSER=aceberg

mod:
	cd src && \
	rm go.mod || true && \
	go mod init $(PKG_NAME) && \
	go mod tidy

run:
	cd src && \
	go run .

fmt:
	cd src && \
	go fmt ./...

go-build:
	cd src && \
	CGO_ENABLED=0 go build -o build-artifacts-1667922535/go-LinkSaver .

docker-build:
	docker build -t $(DUSER)/$(PKG_NAME) .

docker-run:
	docker rm $(PKG_NAME) || true
	docker run --name $(PKG_NAME) \
	-e "TZ=Asia/Novosibirsk" \
	-v ~/.dockerdata/$(PKG_NAME):/data/$(PKG_NAME) \
	-p 8841:8841 \
	$(DUSER)/$(PKG_NAME)