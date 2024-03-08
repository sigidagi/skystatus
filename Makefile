.PHONY: build clean test serve run-compose-test
VERSION := $(shell git describe --always |sed -e "s/^v//")

EXEC_NAME := skystatus

build:
	go build $(GO_EXTRA_BUILD_ARGS) -ldflags "-s -w -X main.version=$(VERSION)" -o build/${EXEC_NAME} cmd/skystatus/main.go

clean:
	@echo "Cleaning up workspace"
	@rm -rf build
	@rm -rf dist

test:
	@echo "Running tests"
	@rm -f coverage.out
	@golint ./...
	@go vet ./...
	@go test -cover -v -coverprofile coverage.out -p 1 ./...

snapshot:
	@goreleaser --snapshot

dev-requirements:
	go install golang.org/x/lint/golint
	go install github.com/goreleaser/goreleaser
	go install github.com/goreleaser/nfpm

serve: build
	./build/${EXEC_NAME}

#deploy:
	#@ssh matter-demo 'systemctl stop ${EXEC_NAME}.service'
	#@echo "Stoped ${EXEC_NAME} service on 'matter-demo'"
	#@scp ./build/${EXEC_NAME_RPI} matter-demo:/usr/bin
	#@scp ./packaging/${EXEC_NAME}.service matter-demo:/etc/systemd/system
	#@scp ./packaging/${EXEC_NAME}.toml matter-demo:/etc/${EXEC_NAME}
	#@ssh matter-demo 'sudo systemctl daemon-reload'
	#@ssh matter-demo 'sudo systemctl start ${EXEC_NAME}.service'
	#@echo "Deplayment Done!"



