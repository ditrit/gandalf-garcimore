package test

import (
	"fmt"
	"strconv"
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
	fmt.Println(command)
	//id := r.client.CreateIteratorEvent()
	//event := r.client.WaitEvent("test", "test", id)
	for i := 1; i < 5; i++ {
		fmt.Println(i)
		r.client.SendEvent(command.GetUUID(), "10000", strconv.Itoa(i*20), "test")
		time.Sleep(time.Duration(1) * time.Millisecond)

	}
	r.client.SendEvent(command.GetUUID(), "10000", "SUCCES", "test")

}
