package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"time"

	"microservices-backend-go/backend"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	_ "github.com/polarismesh/grpc-go-polaris"
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

	for i := 0; i < 4; i++ {
		start := time.Now().Unix()
		resp, err := dbProxy.DescribeUserList(ctx, &backend.DescribeUserListRequest{})
		log.Printf("send message, resp (%v), err(%v)", resp, err)

		bys, _ := json.Marshal(resp.GetUserList())
		fmt.Println(string(bys))
		fmt.Println(time.Now().Unix() - start)
	}
}

func compute(a int, b int) int {

	return a + b
}
