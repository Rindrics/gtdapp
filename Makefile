.DEFAULT_GOAL := help
help: FORCE		## Show this help
	@grep -hE '^[0-9a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) \
	    | sed -e 's/FORCE/     /' -e 's/##//'

.PHONY: run
run: src/gtdapp FORCE  ## Run gtdapp
	export FLASK_ENV=development &&\
	export FLASK_APP=src/gtdapp/app &&\
	flask run

.PHONY: FORCE
FORCE:
