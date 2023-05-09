PROTODIR=proto
OUTDIR=api/gen/api

.PHONY:gen
gen:
	protoc -I$(PROTODIR) \
	--go_out=$(OUTDIR) --go_opt=paths=source_relative \
	--go-grpc_out=$(OUTDIR) --go-grpc_opt=paths=source_relative $(PROTODIR)/*.proto

