COMMON_PROTODIR=../../common/proto
VERSION := $$(git describe --tags --always --dirty)
BRANCH := $$(git symbolic-ref --short HEAD)

.PHONY: gen
gen:
	out_dir=internal; \
	protoc -Iproto -I$(COMMON_PROTODIR) \
	--go_out=$$out_dir --go_opt=paths=source_relative \
	--go-grpc_out=$$out_dir \
	--go-grpc_opt=paths=source_relative \
	--grpc-gateway_out=$$out_dir \
	--grpc-gateway_opt logtostderr=true \
	--grpc-gateway_opt paths=source_relative \
	--openapiv2_out=$$out_dir \
	--openapiv2_opt logtostderr=true \
	proto/*.proto; \
	protoc -I$(COMMON_PROTODIR) \
	--go_out=$$out_dir --go_opt=paths=source_relative \
	--go-grpc_out=$$out_dir \
	--go-grpc_opt=paths=source_relative \
	$(COMMON_PROTODIR)/*.proto

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
