PROTODIR=./proto
COMMONDIR=../../proto
GOOGLE_APIS=$(COMMONDIR)/google/api/
VERSION := $$(git describe --tags --always --dirty)
BRANCH := $$(git symbolic-ref --short HEAD)

.PHONY: gen
gen:
	out_dir=internal; \
	protoc -I$(COMMONDIR) -I$(PROTODIR) \
	--go_out=$$out_dir --go_opt=paths=source_relative \
	--go-grpc_out=$$out_dir \
	--go-grpc_opt=paths=source_relative \
	$(COMMONDIR)/common.proto \
	$(PROTODIR)/inbox.proto;

.PHONY: build
build: cmd internal
	Docker build \
		--platform linux/amd64 \
		-t rindrics/gtdapp-inbox:${BRANCH} \
		-t rindrics/gtdapp-inbox:latest \
		-t rindrics/gtdapp-inbox:${VERSION} \
		.

.PHONY: push
push: build
	Docker push rindrics/gtdapp-inbox:${BRANCH}
	Docker push rindrics/gtdapp-inbox:latest
	Docker push rindrics/gtdapp-inbox:${VERSION}

.PHONY: test
test: cmd internal test/
	go test ./test/...

.PHONY: run
run: cmd internal
	go run ./cmd

.PHONY: ls
ls:
	grpc_cli ls localhost:50051 gtd.inbox.Inbox -l

.PHONY: call
call:
	grpc_cli call localhost:50051 gtd.inbox.Inbox.Collect \
	'title: "Hello, world" description: "Greet to the world"'
