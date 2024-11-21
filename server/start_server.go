package server

import (
	"context"
	"net"

	lo "main/libreoffice"
	proto "main/proto/gobre"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type GobreServer struct {
	proto.UnimplementedGobreServer
}

func (s GobreServer)HandleRequest(
	ctx context.Context, 
	param *proto.Request,
)(*proto.Response, error){
	fileData := lo.HandleConvertFile(
		param.OriginalFileType,
		param.NewFileType,
		param.FileData,
	)	

	return &proto.Response{FileData: fileData}, nil
}

func StartServer(ctx context.Context) {
	listener, listenErr := net.Listen("tcp", ":8081")
	if listenErr != nil {
		panic(listenErr)
	}

	server := grpc.NewServer()
	reflection.Register(server) //Enabled for clients that support reflection
	proto.RegisterGobreServer(server, GobreServer{})

	// Start the server in a separate goroutine
	go func() {
		<-ctx.Done()
		server.Stop()
	}()

	// Serve the server
	server.Serve(listener)
}
