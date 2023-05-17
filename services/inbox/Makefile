COMMON_PROTODIR=../../common/proto

.PHONY: gen
gen:
	out_dir=internal; \
	protoc -Iproto -I$(COMMON_PROTODIR) \
	--go_out=$$out_dir --go_opt=paths=source_relative \
	--go-grpc_out=$$out_dir \
	--go-grpc_opt=paths=source_relative \
	proto/*.proto; \
	protoc -I$(COMMON_PROTODIR) \
	--go_out=$$out_dir --go_opt=paths=source_relative \
	--go-grpc_out=$$out_dir \
	--go-grpc_opt=paths=source_relative \
	$(COMMON_PROTODIR)/*.proto

.PHONY: build
build: cmd internal
	go build -o app ./cmd

.PHONY: test
test: cmd internal test/
	go test ./...

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
