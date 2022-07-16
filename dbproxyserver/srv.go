package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"net"

	polaris "github.com/polarismesh/grpc-go-polaris"
	"microservices-backend-go/backend"
)

type server struct {
	backend.UnimplementedDBProxySrvServer
}

func (s *server) DescribeUserList(context.Context, *backend.DescribeUserListRequest) (*backend.DescribeUserListResponse, error) {
	return &backend.DescribeUserListResponse{
		UserList: []*backend.User{
			{
				Name: "Jack",
			},
		},
	}, nil
}

const (
	listenPort = 16010
)

func runGRpcSrv() {
	srv := grpc.NewServer()
	backend.RegisterDBProxySrvServer(srv, &server{})
	address := fmt.Sprintf("0.0.0.0:%d", listenPort)
	listen, err := net.Listen("tcp", address)
	if err != nil {
		panic(err)
	}

	_, err = polaris.Register(srv, listen)

	if err != nil {
		panic(err)
	}

	//go func() {
	//	c := make(chan os.Signal, 1)
	//	signal.Notify(c)
	//	s := <-c
	//	log.Printf("receive quit signal: %v", s)
	//	// 执行北极星的反注册命令
	//	pSrv.Deregister()
	//	srv.GracefulStop()
	//}()
	err = srv.Serve(listen)
	if nil != err {
		log.Printf("listen err: %v", err)
	}
}
