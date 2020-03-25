package test

type WorkerSend struct {
	client *ClientGrpcTest
}

func NewWorkerSend(identity, connection string) *WorkerSend {
	workerSend := new(WorkerSend)
	workerSend.client = NewClientGrpcTest(identity, connection)

	return workerSend
}

func (r WorkerSend) Run() {
	//r.ClientCommand.SendCommand()
	r.client.SendEvent("test", "100000000", "test", "test", "test")
}
