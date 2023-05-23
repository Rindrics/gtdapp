PROTODIR=../../proto/
GOOGLE_APIS=$(PROTODIR)google/api/
VERSION := $$(git describe --tags --always --dirty)
BRANCH := $$(git symbolic-ref --short HEAD)

.PHONY: gen
gen:
	out_dir=internal; \
	protoc -I$(PROTODIR) -I$(GOOGLE_APIS) \
	--go_out=$$out_dir --go_opt=paths=source_relative \
	--go-grpc_out=$$out_dir \
	--go-grpc_opt=paths=source_relative \
	--grpc-gateway_out=$$out_dir \
	--grpc-gateway_opt logtostderr=true \
	--grpc-gateway_opt paths=source_relative \
	--openapiv2_out=$$out_dir \
	--openapiv2_opt logtostderr=true \
	$(PROTODIR)*.proto;
