package test

import (
	"fmt"
	"time"
)

type WorkerReceive struct {
	client *ClientGrpcTest
}

func NewWorkerReceive(identity, connection string) *WorkerReceive {
	workerReceive := new(WorkerReceive)
	workerReceive.client = NewClientGrpcTest(identity, connection)

	return workerReceive
}

func (r WorkerReceive) Run() {
	id := r.client.CreateIteratorCommand()
	command := r.client.WaitCommand("test", id)
	//id := r.client.CreateIteratorEvent()
	//event := r.client.WaitEvent("test", "test", id)
	//SLEEP
	for i := 1; i < 5; i++ {
		r.client.SendEvent(command.GetUUID(), "10000", string(i*20), "test")
		time.Sleep(time.Duration(10) * time.Millisecond)

	}
	r.client.SendEvent(command.GetUUID(), "10000", "test", "SUCCES")

	fmt.Println(command)
}
