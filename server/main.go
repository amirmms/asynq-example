package main

import (
	"log"

	"asynq/task"
	"github.com/hibiken/asynq"
)

func main() {
	srv := asynq.NewServer(asynq.RedisClientOpt{Addr: "127.0.0.1:6379"}, asynq.Config{
		Concurrency: 1,
		Queues: map[string]int{
			"default": 1,
		}})

	mux := asynq.NewServeMux()

	mux.HandleFunc(task.SMSSendTag, task.HandleSMSSendTask)

	if err := srv.Run(mux); err != nil {
		log.Fatalf("could not run server: %v", err)
	}
}
