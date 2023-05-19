COMMON_PROTODIR=common/proto
PROTODIR=services

# Language settings
GO_SERVICES=services/inbox

.PHONY: gen
gen:
	for dir in $(PROTODIR)/*/proto; do \
		if [ -d $$dir ]; then \
			base_dir=$$(dirname $$dir); \
			out_dir=$$base_dir; \
			lang_opt="--go_out=$$out_dir --go_opt=paths=source_relative --go-grpc_out=$$out_dir --go-grpc_opt=paths=source_relative"; \
			for go_service in $(GO_SERVICES); do \
				if [ $$base_dir = $$go_service ]; then \
					out_dir=$$base_dir/internal; \
					lang_opt="--go_out=$$out_dir --go_opt=paths=source_relative --go-grpc_out=$$out_dir --go-grpc_opt=paths=source_relative"; \
				fi; \
			done; \
			protoc -I$$dir -I$(COMMON_PROTODIR) \
			$$lang_opt \
			$$dir/*.proto; \
		fi; \
	done

# -------
INFRA=tffile/environment/infra
OPERATIONS = plan apply destroy
ENVIRONMENTS = dev prd

define rule_template
.PHONY: infra-$1-$2
infra-$1-$2: $(INFRA)
	./script/$1_infra.sh infra $2
endef
$(foreach op,$(OPERATIONS),$(foreach env,$(ENVIRONMENTS),$(eval $(call rule_template,$(op),$(env)))))
