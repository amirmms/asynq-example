package task

import (
	"context"
	"encoding/json"
	"log"

	"github.com/hibiken/asynq"
)

const SMSSendTag = "sms:send"

type SMSSendPayload struct {
	Mobile  string
	Message string
}

func (p SMSSendPayload) NewSMSSendTask() (*asynq.Task, error) {
	jsonPayload, err := json.Marshal(p)
	if err != nil {
		return nil, err
	}

	task := asynq.NewTask(SMSSendTag, jsonPayload)

	return task, nil
}

func HandleSMSSendTask(ctx context.Context, task *asynq.Task) error {
	var payload SMSSendPayload
	if err := json.Unmarshal(task.Payload(), &payload); err != nil {
		return err
	}

	log.Printf("Send Message To Mobile Number : %s , Message Is : %s \n", payload.Mobile, payload.Message)

	return nil
}
