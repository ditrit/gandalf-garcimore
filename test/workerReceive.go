package test

import "fmt"

type WorkerReceive struct {
	client *ClientGrpcTest
}

func NewWorkerReceive(identity, connection string) *WorkerReceive {
	workerReceive := new(WorkerReceive)
	workerReceive.client = NewClientGrpcTest(identity, connection)

	return workerReceive
}

func (r WorkerReceive) Run() {
	//r.ClientCommand.SendCommand()
	id := r.client.CreateIteratorEvent()
	event := r.client.WaitEvent("test", "test", id)
	fmt.Println(event)
}
