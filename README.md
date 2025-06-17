# rss-consumer
Simple RSS consumer AWS Lambda function.

## Dependencies

* [AWS SAM CLI](https://docs.aws.amazon.com/serverless-application-model/latest/developerguide/install-sam-cli.html)

* [Docker Desktop](https://docs.docker.com/desktop/)

* [go1.24+](https://go.dev/dl/)

## Getting Started

The function `RSSConsumerFunction` will run via EventBridge schedule every **1 day** and retrieve all items in the feed defined in `RSS_URL` over the past `OFFSET_DAYS`. You can modify the behavior of the function by updating the values accordingly.

## Local Development

You can build and invoke the function locally using the following command:

```bash
make invoke
```

By default, it will take a sample event from [event.json](./events/event.json) and retrieve the RSS titles released within 1 day of `2025-05-19T00:00:00Z` (the specified event time). You can modify the event time and/or the `OFFSET_DAYS` variable to alter the behavior of the function.

## Deploying to AWS

You can deploy the function and associated EventBridge schedule using the following command:

```bash
make deploy
```

It will invoke the function once per day.