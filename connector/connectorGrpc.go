package connector

import (
	"context"
	"core/message"
	"fmt"
	"log"
	"net"
)

type ConnectorGrpc struct {
	connectorGrpcConnection string
	connectorGrpcServer     *grpc.Server
	lis                     *net.Listener
}

func NewConnectorGrpc(connectorGrpcConnection string) ConnectorGrpc {
	connectorGrpc = new(ConnectorGrpc)
	connectorGrpc.lis, err := net.Listen("tcp", r.ConnectorCommandWorkerConnection)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	connectorGrpc.connectorGrpcServer = grpc.NewServer()

	pb.RegisterConnectorCommandServer(connectorGrpc.connectorGrpcServer, &connectorGrpc)
	pb.RegisterConnectorEventServer(connectorGrpc.connectorGrpcServer, &connectorGrpc)
}

//GRPC
//startGrpcServer :
func (r ConnectorGrpc) startGrpcServer() {

	r.ConnectorCommandGrpcServer.Serve(lis)
}

//SendCommandMessage :
func (r ConnectorGrpc) SendCommandMessage(ctx context.Context, in *pb.CommandMessage) (*pb.CommandMessageUUID, error) {
	commandMessage := message.CommandMessageFromGrpc(in)

	go commandMessage.SendMessageWith(r.ConnectorCommandSendToAggregator)

	return &pb.CommandMessageUUID{UUID: commandMessage.UUID}, nil
}

//SendCommandMessageReply :
func (r ConnectorGrpc) SendCommandMessageReply(ctx context.Context, in *pb.CommandMessageReply) (*pb.Empty, error) {
	commandMessageReply := message.CommandMessageReplyFromGrpc(in)

	go commandMessageReply.SendMessageWith(r.ConnectorCommandSendToAggregator)

	return &pb.Empty{}, nil
}

//WaitCommandMessage :
func (r ConnectorGrpc) WaitCommandMessage(ctx context.Context, in *pb.CommandMessageWait) (commandMessage *pb.CommandMessage, err error) {
	target := in.GetWorkerSource()
	iterator := NewIterator(r.ConnectorMapCommandNameCommandMessage)

	r.ConnectorMapWorkerIterators[target] = append(r.ConnectorMapWorkerIterators[target], iterator)

	go r.runIteratorCommandMessage(target, in.GetValue(), iterator, r.ConnectorCommandChannel)
	messageChannel := <-r.ConnectorCommandChannel
	commandMessage = message.CommandMessageToGrpc(messageChannel)

	return
}

//WaitCommandMessageReply :
func (r ConnectorGrpc) WaitCommandMessageReply(ctx context.Context, in *pb.CommandMessageWait) (commandMessageReply *pb.CommandMessageReply, err error) {
	target := in.GetWorkerSource()
	iterator := NewIterator(r.ConnectorMapUUIDCommandMessageReply)

	r.ConnectorMapWorkerIterators[target] = append(r.ConnectorMapWorkerIterators[target], iterator)

	go r.runIteratorCommandMessageReply(target, in.GetValue(), iterator, r.ConnectorCommandReplyChannel)
	messageReplyChannel := <-r.ConnectorCommandReplyChannel
	commandMessageReply = message.CommandMessageReplyToGrpc(messageReplyChannel)

	return
}

//SendEventMessage :
//TODO REVOIR SERVICE
func (r ConnectorGrpc) SendEventMessage(ctx context.Context, in *pb.EventMessage) (*pb.Empty, error) {
	eventMessage := message.EventMessageFromGrpc(in)
	fmt.Println(eventMessage)

	go eventMessage.SendMessageWith(r.ConnectorEventSendToAggregator)

	return &pb.Empty{}, nil
}

//WaitEventMessage :
func (r ConnectorGrpc) WaitEventMessage(ctx context.Context, in *pb.EventMessageWait) (messageEvent *pb.EventMessage, err error) {
	target := in.GetWorkerSource()
	iterator := NewIterator(r.ConnectorMapEventNameEventMessage)

	r.ConnectorMapWorkerIterators[in.GetEvent()] = append(r.ConnectorMapWorkerIterators[in.GetEvent()], iterator)

	go r.runIterator(target, in.GetEvent(), iterator, r.ConnectorEventChannel)

	messageChannel := <-r.ConnectorEventChannel
	messageEvent = message.EventMessageToGrpc(messageChannel)

	return
}
