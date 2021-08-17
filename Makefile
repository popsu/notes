.DEFAULT_GOAL := help

.PHONY: help
help: ## This help
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

run-mkdocs: ## Serve mkdocs locally
	xdg-open http://localhost:8000/notes/index.html
	mkdocs serve -f ./mkdocs.yml -a 0.0.0.0:8000
