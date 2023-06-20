PROTODIR=./proto
COMMONDIR=../../proto
VERSION := $$(git describe --tags --always --dirty)
BRANCH := $$(git symbolic-ref --short HEAD)

.PHONY: gen
gen:
	grpc_tools_node_protoc \
	    -I$(COMMONDIR) -I$(PROTODIR) \
	    --plugin=protoc-gen-ts=`pwd`/node_modules/.bin/protoc-gen-ts \
	    --ts_out=service=grpc-web:./src/generated \
	    $(PROTODIR)/projectmanager.proto

.PHONY: test
test:
	@echo "Running tests..."
	@npm test

.PHONY: build
build:
	@echo "Building..."
	@npm run build
