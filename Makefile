VERSION := 0.0.1

DOCKER_REG = bamdockerhub
DOCKER_IMAGE = project-template
DOCKER_IMAGE_TAG = $(VERSION)
USER = $(shell whoami)

FILES=./pkg/contracts/*

docker-build:
	docker build -t $(DOCKER_REG)/$(DOCKER_IMAGE):$(DOCKER_IMAGE_TAG) .

docker-push:
	docker push $(DOCKER_REG)/$(DOCKER_IMAGE):$(DOCKER_IMAGE_TAG)

docker-build-dev:
	docker build -t $(DOCKER_REG)/$(DOCKER_IMAGE):$(USER) .

docker-push-dev:
	docker push $(DOCKER_REG)/$(DOCKER_IMAGE):$(USER)

install:
	go install -v

serve:
	cd cmd/serve; go run main.go

build:
	go build -v ./...

deps:
	go get -v ./...

test:
	go test ./... -cover

lint: 
	gofmt -s -w .
	gofmt -r '(a) -> a' -l *.go .
	gofmt -r '(a) -> a' -w *.go .
	gofmt -r 'α[β:len(α)] -> α[β:]' -w $GOROOT/src .

get_tools:
	go get -u github.com/ethereum/go-ethereum
	cd $(GOPATH)/src/github.com/ethereum/go-ethereum
	make
	make devtools
	go get -u github.com/blockchain-abstraction-middleware/abi-extractor

transpile_contracts:
	for file in $(FILES); do \
		FILENAME=$$(echo $$file | cut -d '/' -f4 | cut -d '.' -f1); \
		FILEPATH=$$(echo $$file | cut -d '.' -f2); \
		abi-extractor $$file; \
		mkdir ./$$FILEPATH ; \
		abigen --abi=.$$FILEPATH.abi --pkg=$$FILENAME --out=.$$FILEPATH/$$FILENAME.go; \
	done

.PHONY: \
	install \
	build \
	serve \
	deps \
	lint \
	test \