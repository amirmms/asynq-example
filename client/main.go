package main

import (
	"fmt"
	"log"

	"asynq/task"
	"github.com/hibiken/asynq"
)

func main() {
	client := asynq.NewClient(asynq.RedisClientOpt{Addr: "127.0.0.1:6379"})
	defer client.Close()

	for i := 0; i < 100; i++ {
		t, err := createNewSMSSendTask("09131234567", "hi this is a test message")
		if err != nil {
			log.Fatal(err)
		}

		enqueue, err := client.Enqueue(t)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println(enqueue.ID)
	}
}

func createNewSMSSendTask(mobile, message string) (*asynq.Task, error) {
	s := task.SMSSendPayload{
		Mobile:  mobile,
		Message: message,
	}

	return s.NewSMSSendTask()
}
