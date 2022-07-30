protoc -I  dbproxyserver/proto \
    -I /home/voyagerma/go/pkg/mod/github.com/grpc-ecosystem/grpc-gateway@v1.16.0/third_party/googleapis \
    --go_out ./backend/  \
    --go-grpc_out ./backend/   \
    dbproxyserver/proto/srv.proto


# generate pb file
protoc -I  dbproxyserver/proto \
    -I /home/voyagerma/go/pkg/mod/github.com/grpc-ecosystem/grpc-gateway@v1.16.0/third_party/googleapis \
    --include_imports --descriptor_set_out=proto.pb \
     dbproxyserver/proto/srv.proto
