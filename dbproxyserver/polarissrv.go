package main

import (
	"flag"
	"fmt"
	"log"
	"microservices-backend-go/dbproxyserver/config"
	"net"

	"microservices-backend-go/backend"
	"microservices-backend-go/dbproxyserver/api"

	polaris "github.com/polarismesh/grpc-go-polaris"

	"google.golang.org/grpc"
)

func startRPCSrvWithPolarisMesh() {
	flag.Parse()

	config.ReadConfigFile()

	srv := grpc.NewServer()
	backend.RegisterDBProxySrvServer(srv, &api.DBProxySrvImp{})
	address := fmt.Sprintf("0.0.0.0:%d", config.listenPort)
	listen, err := net.Listen("tcp", address)
	if err != nil {
		panic(err)
	}
	_, err = polaris.Register(srv, listen)
	if err != nil {
		panic(err)
	}
	err = srv.Serve(listen)
	if nil != err {
		log.Printf("listen err: %v", err)
	}
}
