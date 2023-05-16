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
