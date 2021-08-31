BUILD_ENVPARMS:=CGO_ENABLED=0
BUILD_ENVPARMS_LINUX:=CGO_ENABLED=0 GOOS=linux GOARCH=amd64
BIN?=./bin/law-site-server
BINLINUX?=./bin/law-site-server-linux

build:
	$(info #Building...)
	$(BUILD_ENVPARMS) go build -o $(BIN) ./cmd/law-site-server

build_linux:
	$(info #Building...)
	$(BUILD_ENVPARMS_LINUX) go build -o $(BINLINUX) ./cmd/law-site-server

start:
	$(info #Starting...)
	$(BUILD_ENVPARMS) go run -ldflags "$(LDFLAGS)" ./cmd/law-site-server --port=8888 --host=0.0.0.0

run:
	$(info #Running...)
	./bin/law-site-server --port=8888

run_linux:
	$(info #Running...)
	./bin/law-site-server-linux --port=8888 --host=0.0.0.0

run_linux_bg:
	$(info #Running...)
	nohup ./bin/law-site-server-linux --port=8888 --host=0.0.0.0 &