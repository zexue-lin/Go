package main

import (
	"io"
	"net/http"
	"net/rpc"
	"net/rpc/jsonrpc"
)

type HelloService struct{}

func (s *HelloService) Hello(request string, reply *string) error {
	*reply = "Hello, " + request + "!"
	return nil
}
func main() {
	_ = rpc.RegisterName("HelloService", &HelloService{})

	http.HandleFunc("/jsonrpc", func(w http.ResponseWriter, r *http.Request) {

		var conn io.ReadWriteCloser = struct {
			io.Writer
			io.ReadCloser
		}{
			ReadCloser: r.Body,
			Writer:     w,
		}
		rpc.ServeRequest(jsonrpc.NewServerCodec(conn))
	})
	http.ListenAndServe(":1234", nil)
}
