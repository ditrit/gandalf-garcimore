package test

import (
	"fmt"
	"strconv"
	"time"
)

type WorkerCliReceive struct {
	client      *ClientGrpcTest
	messageType string
	value       string
	topic       string
}

func NewWorkerCliReceive(identity, connection, messageType, value, topic string) *WorkerCliReceive {
	workerCliReceive := new(WorkerCliReceive)
	workerCliReceive.messageType = messageType
	workerCliReceive.value = value
	workerCliReceive.topic = topic
	workerCliReceive.client = NewClientGrpcTest(identity, connection)

	return workerCliReceive
}

func (r WorkerCliReceive) Run() {
	if r.messageType == "cmd" {
		id := r.client.CreateIteratorCommand()
		command := r.client.WaitCommand(r.value, id)
		fmt.Println(command)
		//id := r.client.CreateIteratorEvent()
		//event := r.client.WaitEvent("test", "test", id)
		for i := 1; i < 5; i++ {
			r.client.SendEvent(command.GetUUID(), "10000", strconv.Itoa(i*20), "test")
			time.Sleep(time.Duration(1) * time.Millisecond)

		}
		r.client.SendEvent(command.GetUUID(), "10000", "SUCCES", "test")

	} else if r.messageType == "evt" {
		id := r.client.CreateIteratorEvent()
		event := r.client.WaitEvent(r.value, r.topic, id)
		fmt.Println(event)
	}

}
