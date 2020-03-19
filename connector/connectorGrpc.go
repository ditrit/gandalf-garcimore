package connector

import (
	"context"
	pb "garcimore/grpc"
	"log"
	"net"
	"shoset/msg"
	sn "shoset/net"
	"time"

	"google.golang.org/grpc"
)

type ConnectorGrpc struct {
	GrpcConnection     string
	ShosetConn         *sn.ShosetConn
	MapWorkerIterators map[string][]*msg.Iterator
	CommandChannel     chan msg.Message
	EventChannel       chan msg.Message
}

func NewConnectorGrpc(GrpcConnection string, shosetConn *sn.ShosetConn) (connectorGrpc ConnectorGrpc, err error) {
	connectorGrpc.ShosetConn = shosetConn
	connectorGrpc.GrpcConnection = GrpcConnection
	connectorGrpc.MapWorkerIterators = make(map[string][]*msg.Iterator)
	connectorGrpc.CommandChannel = make(chan msg.Message)
	connectorGrpc.EventChannel = make(chan msg.Message)

	return
}

//GRPC
//startGrpcServer :
func (r ConnectorGrpc) startGrpcServer() {
	lis, err := net.Listen("tcp", r.GrpcConnection)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	connectorGrpcServer := grpc.NewServer()

	pb.RegisterConnectorCommandServer(connectorGrpcServer, &r)
	pb.RegisterConnectorEventServer(connectorGrpcServer, &r)
	connectorGrpcServer.Serve(lis)
}

//SendCommandMessage :
func (r ConnectorGrpc) SendCommandMessage(ctx context.Context, in *pb.CommandMessage) (*pb.CommandMessageUUID, error) {
	cmd := pb.CommandFromGrpc(in)
	ch := r.ShosetConn.GetCh()
	thisOne := ch.GetBindAddr()

	r.ShosetConn.GetCh().ConnsByType.Get("a").Iterate(
		func(key string, val *sn.ShosetConn) {
			if key != r.ShosetConn.GetBindAddr() && key != thisOne {
				val.SendMessage(cmd)
			}
		},
	)

	return &pb.CommandMessageUUID{UUID: cmd.UUID}, nil
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
	target := in.GetWorkerSource()
	ch := r.ShosetConn.GetCh()
	iterator := msg.NewIterator(ch.Queue["cmd"])

	r.MapWorkerIterators[target] = append(r.MapWorkerIterators[target], iterator)

	go r.runIterator(target, in.GetValue(), "cmd", iterator, r.CommandChannel)
	messageChannel := <-r.CommandChannel
	commandMessage = pb.CommandToGrpc(messageChannel.(msg.Command))

	return
}

func (r ConnectorGrpc) runIterator(target, value, msgtype string, iterator *msg.Iterator, channel chan msg.Message) {
	notfound := true
	for notfound {
		iterator.PrintQueue()
		messageIterator := iterator.Get()

		if messageIterator != nil {
			if msgtype == "cmd" {
				message := (messageIterator.GetMessage()).(msg.Command)

				if value == message.Command {
					channel <- message

					notfound = false
				}
			} else if msgtype == "evt" {
				message := (messageIterator.GetMessage()).(msg.Event)

				if value == message.Event {
					channel <- message

					notfound = false
				}
			}

		}

		time.Sleep(time.Duration(2000) * time.Millisecond)
	}
	delete(r.MapWorkerIterators, target)
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
	evt := pb.EventFromGrpc(in)

	ch := r.ShosetConn.GetCh()
	thisOne := ch.GetBindAddr()

	r.ShosetConn.GetCh().ConnsByType.Get("a").Iterate(
		func(key string, val *sn.ShosetConn) {
			if key != r.ShosetConn.GetBindAddr() && key != thisOne {
				val.SendMessage(evt)
			}
		},
	)

	return &pb.Empty{}, nil
	/* 	eventMessage := message.EventMessageFromGrpc(in)
	   	fmt.Println(eventMessage)

	   	go eventMessage.SendMessageWith(r.ConnectorEventSendToAggregator)

	   	return &pb.Empty{}, nil */
}

//WaitEventMessage :
func (r ConnectorGrpc) WaitEventMessage(ctx context.Context, in *pb.EventMessageWait) (messageEvent *pb.EventMessage, err error) {
	target := in.GetWorkerSource()
	ch := r.ShosetConn.GetCh()
	iterator := msg.NewIterator(ch.Queue["evt"])

	r.MapWorkerIterators[in.GetEvent()] = append(r.MapWorkerIterators[in.GetEvent()], iterator)

	go r.runIterator(target, in.GetEvent(), "evt", iterator, r.EventChannel)

	messageChannel := <-r.EventChannel
	messageEvent = pb.EventToGrpc(messageChannel.(msg.Event))

	return
}
