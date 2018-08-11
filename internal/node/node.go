package node

import (
	"context"
	"errors"
	"fmt"
	"net"
	"strings"
	"time"

	"github.com/nemo/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type Node struct {
	// Clients map[string]nodeservice.NodeServiceClient
}

func (n *Node) Ping(ctx context.Context, request *nodeservice.PingRequest) (*nodeservice.PingReply, error) {
	fmt.Println(fmt.Sprintf("Message:\t%s", request.Message))
	return &nodeservice.PingReply{Message: "pong!"}, nil
}

// node is server
func (n *Node) ListenAndServe() {
	// n.Clients = make(map[string]nodeservice.NodeServiceClient)
	go n.listen()
	for {
		// busy loop
	}
}

func (n *Node) listen() {
	listener, err := net.Listen("tcp", ":0")
	if err != nil {
		panic("failed!")
	}
	addr, _ := localIP()
	port := strings.Split(listener.Addr().String(), ":")
	fmt.Println(fmt.Sprintf("Serving at: %s:%s", addr, port[len(port)-1]))

	grpcServer := grpc.NewServer()
	nodeservice.RegisterNodeServiceServer(grpcServer, n)
	reflection.Register(grpcServer)

	if err := grpcServer.Serve(listener); err != nil {
		panic("failed!")
	}
}

// node is client
func (n *Node) Connect(addr string) {
	conn, err := grpc.Dial(addr, grpc.WithInsecure())
	if err != nil {
		panic("failed!")
	}
	defer conn.Close()

	client := nodeservice.NewNodeServiceClient(conn)

	now := time.Now()
	resp, err := client.Ping(context.Background(), &nodeservice.PingRequest{Message: "ping!"})
	if err != nil {
		panic("failed!")
	}
	latency := time.Now().Sub(now)

	fmt.Println(fmt.Sprintf("Message:\t%s\nLatency:\t%dms", resp.Message, latency.Nanoseconds()/int64(1000000)))
}

func localIP() (string, error) {
	ifaces, err := net.Interfaces()
	if err != nil {
		return "", err
	}
	for _, iface := range ifaces {
		if iface.Flags&net.FlagUp == 0 {
			continue // interface down
		}
		if iface.Flags&net.FlagLoopback != 0 {
			continue // loopback interface
		}
		addrs, err := iface.Addrs()
		if err != nil {
			return "", err
		}
		for _, addr := range addrs {
			var ip net.IP
			switch v := addr.(type) {
			case *net.IPNet:
				ip = v.IP
			case *net.IPAddr:
				ip = v.IP
			}
			if ip == nil || ip.IsLoopback() {
				continue
			}
			ip = ip.To4()
			if ip == nil {
				continue // not an ipv4 address
			}
			return ip.String(), nil
		}
	}
	return "", errors.New("failed!")
}
