COMMON_PROTODIR=api/common/proto
PROTODIR=api

.PHONY: gen
gen:
	for dir in $(PROTODIR)/*/proto; do \
		if [ -d $$dir ]; then \
			base_dir=$$(dirname $$dir); \
			protoc -I$$dir -I$(COMMON_PROTODIR) \
			--go_out=$$base_dir --go_opt=paths=source_relative \
			--go-grpc_out=$$base_dir --go-grpc_opt=paths=source_relative $$dir/*.proto; \
		fi; \
	done
