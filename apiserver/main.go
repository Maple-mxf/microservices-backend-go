package main

import (
	"context"
	"encoding/json"
	"fmt"
	_ "github.com/polarismesh/grpc-go-polaris"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"microservices-backend-go/backend"
)

const (
	listenPort = 16011
)

func main() {
	// grpc客户端连接获取
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	conn, err := grpc.DialContext(ctx, "polaris://backend.DBProxySrv",
		grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	dbProxy := backend.NewDBProxySrvClient(conn)

	resp, err := dbProxy.DescribeUserList(ctx, &backend.DescribeUserListRequest{})
	log.Printf("send message, resp (%v), err(%v)", resp, err)

	bys, _ := json.Marshal(resp.GetUserList())
	fmt.Println(string(bys))
}
