.DEFAULT_GOAL := help
help: FORCE		## Show this help
	@grep -hE '^[0-9a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) \
	    | sed -e 's/FORCE/     /' -e 's/##//'

.PHONY: run
run: src/gtdapp FORCE  ## Run gtdapp
	poetry run python3 src/gtdapp/main.py

.PHONY: FORCE
FORCE:
