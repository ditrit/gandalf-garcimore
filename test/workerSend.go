package test

type WorkerSend struct {
	client *ClientGrpcTest
}

func NewWorkerSend(identity, connection string) {
	workerSend := new(WorkerSend)
	workerSend.client = NewClientGrpcTest(identity, connection)
}

func (r WorkerSend) Run() {
	for true {
		//r.ClientCommand.SendCommand()
		r.client.SendEvent("test", "0", "test", "test", "test")
	}
}
