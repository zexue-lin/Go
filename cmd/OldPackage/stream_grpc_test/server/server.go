package main

import (
	"Go/cmd/OldPackage/stream_grpc_test/proto"
	"fmt"
	"google.golang.org/grpc"
	"net"
	"time"
)

const PORT = ":50052"

type server struct {
}

// 实现三个方法
func (s *server) GetStream(req *proto.StreamReqData, res proto.Greeter_GetStreamServer) error {
	i := 0
	for {
		i++
		_ = res.Send(&proto.StreamResData{
			Data: fmt.Sprintf("%v", time.Now().Unix()),
		})
		time.Sleep(1 * time.Second)
		if i > 10 {
			break
		}
	}
	return nil
}

func (s *server) PutStream(cliStr proto.Greeter_PutStreamServer) error {
	return nil
}

func (s *server) AllStream(allStr proto.Greeter_AllStreamServer) error {
	return nil
}

func main() {
	lis, err := net.Listen("tcp", PORT)
	if err != nil {
		panic(err)
	}
	s := grpc.NewServer()
	proto.RegisterGreeterServer(s, &server{})

	err = s.Serve(lis)
	if err != nil {
		panic(err)
	}
}
