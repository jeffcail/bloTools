package tron

import (
	"context"
	"errors"
	"google.golang.org/grpc"
	"log"
	"sync"
)

var networks = []string{
	"grpc.trongrid.io:50051",
	"tron-grpc.publicnode.com:443",
	"52.53.189.99:50051",
	"18.196.99.16:50051",
	"34.253.187.192:50051",
	"18.133.82.227:50051",
	"35.180.51.163:50051",
	"54.252.224.209:50051",
	"18.228.15.36:50051",
	"52.15.93.92:50051",
	"34.220.77.106:50051",
	"15.207.144.3:50051",
	"13.124.62.58:50051",
	"15.222.19.181:50051",
	"18.209.42.127:50051",
	"3.218.137.187:50051",
	"34.237.210.82:50051",
	"47.241.20.47:50051",
	"161.117.85.97:50051",
	"161.117.224.116:50051",
	"161.117.83.38:50051",
}

type node struct {
	network string
}

type nodeManager struct {
	nodes       []node
	currentNode int
	mutex       sync.Mutex
}

func newNodeManager(networks []string) *nodeManager {
	nodes := make([]node, len(networks))
	for i, network := range networks {
		nodes[i] = node{network: network}
	}
	return &nodeManager{
		nodes: nodes,
	}
}

func (nm *nodeManager) getNextNode() (node, error) {
	nm.mutex.Lock()
	defer nm.mutex.Unlock()

	if len(nm.nodes) == 0 {
		return node{}, errors.New("no available node")
	}

	n := nm.nodes[nm.currentNode]
	nm.currentNode = (nm.currentNode + 1) % len(nm.nodes)
	return n, nil
}

type GrpcProxy struct {
	nodeManager *nodeManager
	conn        *grpc.ClientConn
}

func NewGrpcProxy() *GrpcProxy {
	return &GrpcProxy{
		nodeManager: newNodeManager(networks),
	}
}

func (proxy *GrpcProxy) connect(ctx context.Context) error {
	node, err := proxy.nodeManager.getNextNode()
	if err != nil {
		return err
	}

	conn, err := grpc.DialContext(ctx, node.network, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		return err
	}

	proxy.conn = conn
	return nil
}

func (proxy *GrpcProxy) Invoke(ctx context.Context, method string, args any, reply any, opts ...grpc.CallOption) error {
	var err error

	for i := 0; i < len(proxy.nodeManager.nodes); i++ {
		// Attempt to connect to the node.
		if proxy.conn == nil {
			if err = proxy.connect(ctx); err != nil {
				log.Printf("Failed to connect to node: %v", err)
				continue
			}
		}

		// Perform the gRPC call.
		err = proxy.conn.Invoke(ctx, method, args, reply, opts...)
		if err == nil {
			return nil
		}

		// If the call failed, reset the connection and try the next node.
		log.Printf("gRPC call failed: %v. Trying next node...", err)
		proxy.conn.Close()
		proxy.conn = nil
	}

	return errors.New("all gRPC nodes failed")
}

func (proxy *GrpcProxy) Close() {
	if proxy.conn != nil {
		proxy.conn.Close()
	}
}
