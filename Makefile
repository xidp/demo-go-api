BUILD_FOLDER = build
BUILD_BINARY = server
# BUILD_BINARY = $(shell cat go.mod | grep module | cut -d " " -f 2 | rev | cut -d "/" -f 1 | rev)

all: b 
# all: build run

# go build -a -v -buildmode='exe' -buildvcs='true' -mod='vendor' -trimpath='true' -compiler='gc' -o ${BUILD_FOLDER}/${BUILD_BINARY} main.go
.PHONY: build
build:
	GOARCH=arm64 GOOS=darwin go build -o ${BUILD_FOLDER}/latest/${BUILD_BINARY}-macOS-arm64 main.go
	GOARCH=amd64 GOOS=darwin go build -o ${BUILD_FOLDER}/latest/${BUILD_BINARY}-macOS-amd64 main.go
	GOARCH=amd64 GOOS=linux go build -o ${BUILD_FOLDER}/latest/${BUILD_BINARY}-linux-amd64 main.go
	GOARCH=amd64 GOOS=windows go build -o ${BUILD_FOLDER}/latest/${BUILD_BINARY}-windows-amd64.exe main.go

.PHONY: b
b:
	go build -o ${BUILD_FOLDER}/${BUILD_BINARY}
	cp ${BUILD_FOLDER}/${BUILD_BINARY} /usr/local/bin

.PHONY: run
run:
	go run main.go
# run: build
# 	./${BUILD_FOLDER}/${BUILD_BINARY}

# dev:
# 	~/go/bin/CompileDaemon -build="go build -o ${BUILD_FOLDER}/${BUILD_BINARY}" -command="./${BUILD_FOLDER}/${BUILD_BINARY}"

.PHONY: test
install:
	$(info ******************** downloading dependencies ********************)
	go get -v ./...

.PHONY: test
test:
	$(info ******************** running tests ********************)
	go test -v ./... -count=1
# count=1 makes it not to cache

.PHONY: test_cov
test_cov:
	mkdir -p coverage
	go test ./... -coverprofile=./coverage/coverage.out

.PHONY: vet
vet:
	go vet

.PHONY: lint
lint:
	golangci-lint run --enable-all
# If you are using go mod, make sure to add "golangci-lint" to your go.mod file.

.PHONY: clean
clean:
	go clean
	go mod tidy
	rm -Rf ./${BUILD_FOLDER} ./coverage
#  rm ${BUILD_BINARY}-darwin
#  rm ${BUILD_BINARY}-linux
#  rm ${BUILD_BINARY}-windows