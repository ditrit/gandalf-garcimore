package test

import "fmt"

type WorkerCliSend struct {
	client      *ClientGrpcTest
	messageType string
	value       string
	topic       string
	payload     string
}

func NewWorkerCliSend(identity, connection, messageType, value, topic, payload string) *WorkerCliSend {
	workerCliSend := new(WorkerCliSend)
	workerCliSend.messageType = messageType
	workerCliSend.value = value
	workerCliSend.topic = topic
	workerCliSend.payload = payload

	workerCliSend.client = NewClientGrpcTest(identity, connection)

	return workerCliSend
}

func (r WorkerCliSend) Run() {
	if r.messageType == "cmd" {
		commandUUID := r.client.SendCommand("100000", "test", r.value, r.payload)
		if commandUUID != nil {
			id := r.client.CreateIteratorEvent()
			for {
				event := r.client.WaitTopic(commandUUID.GetUUID(), id)
				fmt.Println(event)

				if event.GetEvent() == "SUCCES" || event.GetEvent() == "FAIL" {
					fmt.Println(event.GetPayload())
					break
				}
			}
		}
	} else if r.messageType == "evt" {
		r.client.SendEvent(r.topic, "100000", r.value, r.payload)
	}

	//r.client.SendEvent("test", "10000", "test", "test", "test")
}
