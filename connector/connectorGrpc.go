package connector

import (
	"context"
	"fmt"
	pb "garcimore/grpc"
	"garcimore/utils"
	"log"
	"net"
	"shoset/msg"
	sn "shoset/net"
	"time"

	"github.com/google/uuid"
	"google.golang.org/grpc"
)

var sendIndex = 0

type ConnectorGrpc struct {
	GrpcConnection string
	Shoset         sn.Shoset
	//MapWorkerIterators map[string][]*msg.Iterator
	MapIterators   map[string]*msg.Iterator
	CommandChannel chan msg.Message
	EventChannel   chan msg.Message
}

func NewConnectorGrpc(GrpcConnection string, shoset *sn.Shoset) (connectorGrpc ConnectorGrpc, err error) {
	connectorGrpc.Shoset = *shoset
	connectorGrpc.GrpcConnection = GrpcConnection
	//connectorGrpc.MapWorkerIterators = make(map[string][]*msg.Iterator)
	connectorGrpc.MapIterators = make(map[string]*msg.Iterator)
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
	cmd.Tenant = r.Shoset.Context["tenant"].(string)

	ch := r.Shoset
	fmt.Println(r.Shoset.ConnsByAddr)
	r.Shoset.ConnsByAddr.Iterate(
		func(key string, val *sn.ShosetConn) {
			fmt.Println(key)
			fmt.Println(val)
		},
	)
	//thisOne := ch.GetBindAddr()
	fmt.Println("TOTO")
	fmt.Println(ch.ConnsByAddr)
	shosets := utils.GetByType(ch.ConnsByAddr, "a")
	fmt.Println("shosets")
	fmt.Println(shosets)
	index := getSendIndex(shosets)
	fmt.Println("index")
	fmt.Println(index)
	var send = false
	for !send {
		fmt.Println("SEND")
		fmt.Println(shosets[index])
		shosets[index].SendMessage(cmd)
		timeoutSend := time.Duration((int(cmd.GetTimeout()) / len(shosets)))
		time.Sleep(timeoutSend * time.Millisecond)
		fmt.Println("EVT")

		evt := ch.Queue["evt"].GetByUUID(cmd.GetUUID())
		fmt.Println(evt)
		if evt != nil {
			fmt.Println("break")

			break
		}
	}

	/*
		r.Shoset.ConnsByAddr.Iterate(
			func(key string, val *sn.ShosetConn) {
				if key != r.Shoset.GetBindAddr() && key != thisOne && val.ShosetType == "a" {
					val.SendMessage(cmd)
					//WAIT REP
				}
			},
		) */

	return &pb.CommandMessageUUID{UUID: cmd.UUID}, nil
}

//WaitCommandMessage :
func (r ConnectorGrpc) WaitCommandMessage(ctx context.Context, in *pb.CommandMessageWait) (commandMessage *pb.CommandMessage, err error) {
	iterator := msg.NewIterator(r.Shoset.Queue["cmd"])

	//r.MapWorkerIterators[in.GetIteratorId()] = append(r.MapWorkerIterators[in.GetIteratorId()], iterator)
	r.MapIterators[in.GetIteratorId()] = iterator

	go r.runIterator(in.GetIteratorId(), in.GetValue(), "cmd", iterator, r.CommandChannel)
	messageChannel := <-r.CommandChannel
	fmt.Println("TOTOTOTOTOTO")
	commandMessage = pb.CommandToGrpc(messageChannel.(msg.Command))

	return
}

//SendEventMessage :
func (r ConnectorGrpc) SendEventMessage(ctx context.Context, in *pb.EventMessage) (*pb.Empty, error) {
	evt := pb.EventFromGrpc(in)
	evt.Tenant = r.Shoset.Context["tenant"].(string)

	ch := r.Shoset
	thisOne := ch.GetBindAddr()

	r.Shoset.ConnsByAddr.Iterate(
		func(key string, val *sn.ShosetConn) {
			if key != r.Shoset.GetBindAddr() && key != thisOne && val.ShosetType == "a" {
				val.SendMessage(evt)
			}
		},
	)

	return &pb.Empty{}, nil
}

//WaitEventMessage :
func (r ConnectorGrpc) WaitEventMessage(ctx context.Context, in *pb.EventMessageWait) (messageEvent *pb.EventMessage, err error) {
	iterator := msg.NewIterator(r.Shoset.Queue["evt"])
	fmt.Println("WAIT EVENT")
	fmt.Println("QUEUEU")
	fmt.Println(r.Shoset.Queue["cmd"])

	//r.MapWorkerIterators[in.GetIteratorId()] = append(r.MapWorkerIterators[in.GetIteratorId()], iterator)
	r.MapIterators[in.GetIteratorId()] = iterator

	go r.runIterator(in.GetIteratorId(), in.GetEvent(), "evt", iterator, r.EventChannel)

	messageChannel := <-r.EventChannel
	messageEvent = pb.EventToGrpc(messageChannel.(msg.Event))

	return
}

//TODO REFACTORING
//CreateIteratorCommand :
func (r ConnectorGrpc) CreateIteratorCommand(ctx context.Context, in *pb.Empty) (iteratorMessage *pb.IteratorMessage, err error) {
	fmt.Println("CREATE ITERATOR CMD")
	iterator := msg.NewIterator(r.Shoset.Queue["cmd"])
	index := uuid.New()
	//r.MapWorkerIterators[index.String()] = append(r.MapWorkerIterators[index.String()], iterator)
	r.MapIterators[index.String()] = iterator

	iteratorMessage = &pb.IteratorMessage{Id: index.String()}

	return
}

//CreateIteratorEvent :
func (r ConnectorGrpc) CreateIteratorEvent(ctx context.Context, in *pb.Empty) (iteratorMessage *pb.IteratorMessage, err error) {
	fmt.Println("CREATE ITERATOR EVENT")

	iterator := msg.NewIterator(r.Shoset.Queue["evt"])
	index := uuid.New()
	//r.MapWorkerIterators[index.String()] = append(r.MapWorkerIterators[index.String()], iterator)
	r.MapIterators[index.String()] = iterator

	iteratorMessage = &pb.IteratorMessage{Id: index.String()}
	return
}

func (r ConnectorGrpc) runIterator(iteratorId, value, msgtype string, iterator *msg.Iterator, channel chan msg.Message) {

	//fmt.Println("ITERATOR QUEUE")
	//fmt.Println(iterator)
	//iterator.PrintQueue()

	notfound := true
	for notfound {
		fmt.Println("ITERATOR QUEUE")
		iterator.PrintQueue()
		messageIterator := iterator.Get()

		if messageIterator != nil {
			if msgtype == "cmd" {
				fmt.Println("COMMAND")

				message := (messageIterator.GetMessage()).(msg.Command)
				fmt.Println(message)
				fmt.Println(message.GetCommand())
				fmt.Println(message.GetUUID())
				fmt.Println(value)
				if value == message.GetCommand() {
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
	delete(r.MapIterators, iteratorId)
}

func getSendIndex(conns []*sn.ShosetConn) int {
	aux := sendIndex
	sendIndex++
	if sendIndex >= len(conns) {
		sendIndex = 0
	}
	return aux
}
