LAMBDA_NAME=rss_lambda
GOOS=linux
GOARCH=amd64

.PHONY: build clean

build:
	sam build -t template.yml

build-RSSConsumerFunction:
	GOOS=$(GOOS) GOARCH=$(GOARCH) go build -tags lambda.norpc -o bootstrap cmd/main.go
	cp bootstrap $(ARTIFACTS_DIR)/

clean:
	rm -rf .aws-sam

test:
	go test -v ./...

test-coverage:
	go test -v -coverprofile=coverage.out ./...
	go tool cover -html=coverage.out -o coverage.html
	@echo "Coverage report generated: coverage.html"

lint:
	golangci-lint run --timeout 5m

install:
	brew install golangci-lint
	brew upgrade golangci-lint