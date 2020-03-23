package test

import "fmt"

type WorkerReceive struct {
	client *ClientGrpcTest
}

func NewWorkerReceive(identity, connection string) {
	workerReceive := new(WorkerReceive)
	workerReceive.client = NewClientGrpcTest(identity, connection)
}

func (r WorkerReceive) Run() {
	for true {
		//r.ClientCommand.SendCommand()
		id := r.client.CreateIteratorEvent()
		event := r.client.WaitEvent(id, "test", "test")
		fmt.Println(event)
	}
}
