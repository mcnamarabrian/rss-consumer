LAMBDA_LOGICAL_NAME=RSSConsumerFunction
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

deploy:
	@echo "Invoking Lambda function locally..."
	sam deploy --guided --stack-name rss-consumer-stack --capabilities CAPABILITY_IAM

invoke: build
	@echo "Invoking Lambda function locally..."
	sam local invoke $(LAMBDA_NAME) -e events/event.json

install:
	brew install golangci-lint
	brew upgrade golangci-lint

lint:
	golangci-lint run --timeout 5m

test:
	go test -v ./...

test-coverage:
	go test -v -coverprofile=coverage.out ./...
	go tool cover -html=coverage.out -o coverage.html
	@echo "Coverage report generated: coverage.html"
