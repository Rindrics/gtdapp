PROTODIR=./proto
COMMONDIR=../../proto
VERSION := $$(git describe --tags --always --dirty)
BRANCH := $$(git symbolic-ref --short HEAD)

.PHONY: gen
gen:
	grpc_tools_node_protoc \
	    -I$(COMMONDIR) -I$(PROTODIR) \
	    --plugin=protoc-gen-ts=./node_modules/.bin/protoc-gen-ts \
	    --ts_out=service=true:./src/generated \
	    --js_out=import_style=commonjs,binary:./src/generated \
	    --grpc_out=./src/generated \
	    $(COMMONDIR)/google/api/annotations.proto \
	    $(COMMONDIR)/google/api/http.proto \
	    $(PROTODIR)/projectmanager.proto

.PHONY: test
test:
	@echo "Running tests..."
	@npm test

.PHONY: build
build:
	@echo "Building..."
	@npm run build
