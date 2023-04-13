EXPORT_RESULT?=false # for CI please set EXPORT_RESULT to true

#makefile: https://betterprogramming.pub/my-ultimate-makefile-for-golang-projects-fcc8ca20c9bb
#makefile: https://gist.github.com/thomaspoignant/5b72d579bd5f311904d973652180c705

GREEN  := $(shell tput -Txterm setaf 2)
YELLOW := $(shell tput -Txterm setaf 3)
WHITE  := $(shell tput -Txterm setaf 7)
CYAN   := $(shell tput -Txterm setaf 6)
RESET  := $(shell tput -Txterm sgr0)

.PHONY: all test vet fmt lint lint-go build start-listener start-forwarder start-producer-v2

all: test vet fmt lint build

test:
	go test ./...

# TODO Modify to export junit in cicd at some point
# ## Test:
# test: ## Run the tests of the project
# ifeq ($(EXPORT_RESULT), true)
# 	GO111MODULE=off go get -u github.com/jstemmer/go-junit-report
# 	$(eval OUTPUT_OPTIONS = | tee /dev/tty | go-junit-report -set-exit-code > junit-report.xml)
# endif
# 	$(GOTEST) -v -race ./... $(OUTPUT_OPTIONS)

# coverage: ## Run the tests of the project and export the coverage
# 	$(GOTEST) -cover -covermode=count -coverprofile=profile.cov ./...
# 	$(GOCMD) tool cover -func profile.cov
# ifeq ($(EXPORT_RESULT), true)
# 	GO111MODULE=off go get -u github.com/AlekSi/gocov-xml
# 	GO111MODULE=off go get -u github.com/axw/gocov/gocov
# 	gocov convert profile.cov | gocov-xml > coverage.xml
# endif

lint-go: ## Use golintci-lint on your project
	$(eval OUTPUT_OPTIONS = $(shell [ "${EXPORT_RESULT}" == "true" ] && echo "--out-format checkstyle ./... | tee /dev/tty > checkstyle-report.xml" || echo "" ))
	docker run --rm -v $(shell pwd):/app -w /app golangci/golangci-lint:latest-alpine golangci-lint run --deadline=65s $(OUTPUT_OPTIONS)

vet:
	go vet ./...

fmt:
	gofmt -s -w .

lint:
	go list ./... | grep -v /vendor/ | xargs -L1 golint -set_exit_status

build:
	go build -o bins/forwarder ./cmd/snmp-forwarder
	go build -o bins/listener ./cmd/snmp-listener
	go build -o bins/producer_v2 ./cmd/snmp-producer-v2
	go build -o bins/producer_v3 ./cmd/snmp-producer-v3

lint-go: ## Use golintci-lint on your project
	$(eval OUTPUT_OPTIONS = $(shell [ "${EXPORT_RESULT}" == "true" ] && echo "--out-format checkstyle ./... | tee /dev/tty > checkstyle-report.xml" || echo "" ))
	docker run --rm -v $(shell pwd):/app -w /app golangci/golangci-lint:latest-alpine golangci-lint run --deadline=65s $(OUTPUT_OPTIONS)

vendor: ## Copy of all packages needed to support builds and tests in the vendor directory
	go mod vendor

start-listener:
	./bins/listener

start-forwarder:
	./bins/forwarder

start-producer-v2:
	./bins/producer_v2