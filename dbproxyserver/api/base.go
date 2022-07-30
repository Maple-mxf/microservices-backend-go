package api

import (
	"context"

	"microservices-backend-go/backend"
)

func (s *DBProxySrvImp) DescribeUserList(context.Context, *backend.DescribeUserListRequest) (*backend.DescribeUserListResponse, error) {
	return &backend.DescribeUserListResponse{
		UserList: []*backend.User{
			{
				Name: "Jack",
			},
		},
	}, nil
}
