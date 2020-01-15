
GOPATH:=$(shell go env GOPATH)


.PHONY: build
build:
	go build -o ./bin/payment *.go

.PHONY: build_linux
build_linux:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -ldflags '-w' -o ./bin/linux_amd64/payment *.go

.PHONY: test
test:
	go test -v ./... -cover

.PHONY: docker
docker: build_linux
	docker build . -t $(tag)

.PHONY: run
run:
	./bin/payment --registry=$(registry) --transport=$(transport)
