package main

import (
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sqs"
)

func main() {
	args := os.Args
	if len(args) < 2 {
		showUsage(args[0])
		return
	}

	cli, err := createClient()
	if err != nil {
		log.Printf("[ERROR createClient] %s", err.Error())
		return
	}

	list, ok := createQueues(cli, args[1:])
	if !ok {
		return
	}

	log.Printf("[DONE] queues:%d, names:%s", len(list), strings.Join(list, ","))
}

// createClient creates SQS client.
func createClient() (*client, error) {
	sess, err := session.NewSession(awsConfig())
	if err != nil {
		return nil, err
	}

	cli := sqs.New(sess)
	return &client{
		svc: cli,
	}, nil
}

func createQueues(cli *client, list []string) (success []string, ok bool) {
	success = make([]string, 0, len(list))

	for _, q := range list {
		log.Printf("create queue name:%s ...", q)

		has, err := cli.isExist(q)
		switch {
		case err != nil:
			log.Printf("[SQS error] name:%s error:%s", q, err.Error())
			return nil, false
		case has:
			log.Printf("already exists name:%s, skipped", q)
			continue
		}

		err = cli.create(q)
		if err != nil {
			log.Printf("[creation error] name:%s error:%s", q, err.Error())
			continue
		}
		success = append(success, q)

		time.Sleep(100 * time.Millisecond)
	}
	return success, true
}

type client struct {
	svc *sqs.SQS
}

func (c *client) isExist(name string) (bool, error) {
	data, err := c.svc.GetQueueUrl(&sqs.GetQueueUrlInput{
		QueueName: stringPtr(name),
	})

	switch {
	case isNonExistentQueueError(err):
		return false, nil
	case err != nil:
		return false, err
	case data == nil:
		return false, nil
	case *data.QueueUrl != "":
		return true, nil
	default:
		return false, nil
	}
}

func (c *client) create(name string) error {
	_, err := c.svc.CreateQueue(&sqs.CreateQueueInput{
		QueueName: stringPtr(name),
	})
	return err
}

func isNonExistentQueueError(err error) bool {
	const errNonExistentQueue = "NonExistentQueue: "
	if err == nil {
		return false
	}

	return strings.Contains(err.Error(), errNonExistentQueue)
}

func stringPtr(str string) *string {
	return &str
}

func showUsage(str string) {
	fmt.Printf("Specify SQS Queue name to create.\n\n")
	fmt.Printf("Usage:\n")
	fmt.Printf("$ %s <queue names>\n\n", str)
	fmt.Printf("ex) $ %s my_queue1 my_queue2 my_queue3\n", str)
}
