package connector

import (
	"context"
	pb "garcimore/grpc"
	"log"
	"net"
	sn "shoset/net"

	"google.golang.org/grpc"
)

type ConnectorGrpc struct {
	connectorGrpcConnection string
	connectorGrpcServer     *grpc.Server
	lis                     net.Listener
	shosetConn              *sn.ShosetConn
}

func NewConnectorGrpc(connectorGrpcConnection string, shosetConn *sn.ShosetConn) (connectorGrpc ConnectorGrpc, err error) {
	connectorGrpc.shosetConn = shosetConn
	connectorGrpc.lis, err = net.Listen("tcp", connectorGrpcConnection)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	connectorGrpc.connectorGrpcServer = grpc.NewServer()

	pb.RegisterConnectorCommandServer(connectorGrpc.connectorGrpcServer, &connectorGrpc)
	pb.RegisterConnectorEventServer(connectorGrpc.connectorGrpcServer, &connectorGrpc)

	return
}

//GRPC
//startGrpcServer :
func (r ConnectorGrpc) startGrpcServer() {

	r.connectorGrpcServer.Serve(r.lis)
}

//SendCommandMessage :
func (r ConnectorGrpc) SendCommandMessage(ctx context.Context, in *pb.CommandMessage) (*pb.CommandMessageUUID, error) {
	command := pb.CommandFromGrpc(in)
	//TODO
	//r.shosetConn.SendMessage(command)

	return &pb.CommandMessageUUID{UUID: command.UUID}, nil
}

//TODO USELESS
//SendCommandMessageReply :
func (r ConnectorGrpc) SendCommandMessageReply(ctx context.Context, in *pb.CommandMessageReply) (*pb.Empty, error) {
	/* 	commandMessageReply := message.CommandMessageReplyFromGrpc(in)

	   	go commandMessageReply.SendMessageWith(r.ConnectorCommandSendToAggregator)
	*/
	return &pb.Empty{}, nil
}

//WaitCommandMessage :
func (r ConnectorGrpc) WaitCommandMessage(ctx context.Context, in *pb.CommandMessageWait) (commandMessage *pb.CommandMessage, err error) {
	/* 	target := in.GetWorkerSource()
	   	iterator := NewIterator(r.ConnectorMapCommandNameCommandMessage)

	   	r.ConnectorMapWorkerIterators[target] = append(r.ConnectorMapWorkerIterators[target], iterator)

	   	go r.runIteratorCommandMessage(target, in.GetValue(), iterator, r.ConnectorCommandChannel)
	   	messageChannel := <-r.ConnectorCommandChannel
	   	commandMessage = pb.CommandToGrpc(messageChannel) */

	return
}

//TODO USELESS
//WaitCommandMessageReply :
func (r ConnectorGrpc) WaitCommandMessageReply(ctx context.Context, in *pb.CommandMessageWait) (commandMessageReply *pb.CommandMessageReply, err error) {
	/* 	target := in.GetWorkerSource()
	   	iterator := NewIterator(r.ConnectorMapUUIDCommandMessageReply)

	   	r.ConnectorMapWorkerIterators[target] = append(r.ConnectorMapWorkerIterators[target], iterator)

	   	go r.runIteratorCommandMessageReply(target, in.GetValue(), iterator, r.ConnectorCommandReplyChannel)
	   	messageReplyChannel := <-r.ConnectorCommandReplyChannel
	   	commandMessageReply = message.CommandMessageReplyToGrpc(messageReplyChannel)
	*/
	return
}

//SendEventMessage :
func (r ConnectorGrpc) SendEventMessage(ctx context.Context, in *pb.EventMessage) (*pb.Empty, error) {
	event := pb.EventFromGrpc(in)

	//TODO
	//r.shosetConn.SendMessage(event)

	return &pb.Empty{}, nil
	/* 	eventMessage := message.EventMessageFromGrpc(in)
	   	fmt.Println(eventMessage)

	   	go eventMessage.SendMessageWith(r.ConnectorEventSendToAggregator)

	   	return &pb.Empty{}, nil */
}

//WaitEventMessage :
func (r ConnectorGrpc) WaitEventMessage(ctx context.Context, in *pb.EventMessageWait) (messageEvent *pb.EventMessage, err error) {
	/* 	target := in.GetWorkerSource()
	   	iterator := NewIterator(r.ConnectorMapEventNameEventMessage)

	   	r.ConnectorMapWorkerIterators[in.GetEvent()] = append(r.ConnectorMapWorkerIterators[in.GetEvent()], iterator)

	   	go r.runIterator(target, in.GetEvent(), iterator, r.ConnectorEventChannel)

	   	messageChannel := <-r.ConnectorEventChannel
	   	messageEvent = message.EventMessageToGrpc(messageChannel)
	*/
	return
}
