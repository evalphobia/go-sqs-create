# go-sqs-create
SQS queue creation tool

# Install

```bash
$ go get github.com/evalphobia/go-sqs-create
```

# Usage

```bash
$ go-sqs-create <queue names...>

$ go-sqs-create my_queue1 my_queue2 my_queue3
```

Environment parameters are supported:

- Access Key ID:
    - `AWS_ACCESS_KEY_ID`
    - `AWS_ACCESS_KEY`
- Secret Access Key:
    - `AWS_SECRET_ACCESS_KEY`
    - `AWS_SECRET_KEY`
- Region:
    - `AWS_REGION`
- Endpoint:
    - `AWS_SQS_ENDPOINT`

```bash
$ AWS_ACCESS_KEY_ID="XXX" AWS_SECRET_ACCESS_KEY="YYY" AWS_REGION="ap-northeast-1" go-sqs-create my_queue1 my_queue2 my_queue3
```

## local endpoint

Use [Fake SQS](https://github.com/iain/fake_sqs)

```bash
# run fake sqs
$ gem install fake_sqs
$ fake_sqs

# set endpoint to fake sqs
$ AWS_SQS_ENDPOINT="http://localhost:4568" go-sqs-create my_queue1 my_queue2 my_queue3
```
