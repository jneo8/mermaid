TIME=$$(date +%Y-%m-%d_%H:%M)

##@ Show

count-line:  ## Count *.go line in project
	    find . -name '*.go' | xargs wc -l

##@ golang script

go-test:  ## Run go test
	    go test -v --cover ./...

##@ Help

.PHONY: help

help:  ## Display this help
	    @awk 'BEGIN {FS = ":.*##"; printf "\nUsage:\n  make \033[36m<target>\033[0m\n"} /^[a-zA-Z_-0-9]+:.*?##/ { printf "  \033[36m%-15s\033[0m %s\n", $$1, $$2 } /^##@/ { printf "\n\033[1m%s\033[0m\n", substr($$0, 5) } ' $(MAKEFILE_LIST)

.DEFAULT_GOAL := help

