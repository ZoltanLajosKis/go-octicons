PACKAGES  := $(shell go list ./... | grep -v '/vendor/')

.PHONY: generate
generate: vendor
	go run generate/*.go

vendor: | dep
		dep ensure -v

.PHONY: dep
dep:
ifeq ($(shell command -v dep 2> /dev/null),)
		go get -u github.com/golang/dep/cmd/dep
endif

.PHONY: update-deps
update-deps: dep
		dep ensure -update -v
		@touch vendor


.PHONY: test
test: | node npm xmllint
		npm install --no-save octicons@6.0.1
		go test -coverprofile=cover.out -v $(PACKAGES)
		rm -rf node_modules

.PHONY: node
node:
ifeq ($(shell command -v node 2> /dev/null),)
		$(error "node is not available.")
endif

.PHONY: npm
npm:
ifeq ($(shell command -v npm 2> /dev/null),)
		$(error "npm is not available.")
endif

.PHONY: xmllint
xmllint:
ifeq ($(shell command -v xmllint 2> /dev/null),)
		$(error "xmllint is not available.")
endif


.PHONY: clean
clean:
		rm -f octicons.go

.PHONY: clean-all
clean-all: clean
		rm -rf vendor/

