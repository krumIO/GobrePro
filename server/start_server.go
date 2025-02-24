package server

import (
	"context"
	"net"

	lo "main/libreoffice"
	proto "main/proto/gobre"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

  "google.golang.org/grpc/health"
  "google.golang.org/grpc/health/grpc_health_v1"
)

type GobreServer struct {
	proto.UnimplementedGobreServer
}

func (s GobreServer)HandleFileRequest(
	ctx context.Context, 
	param *proto.FileRequest,
)(*proto.FileResponse, error){
	fileData := lo.HandleConvertFile(
		param.OriginalFileType,
		param.NewFileType,
		param.FileData,
	)	

	return &proto.FileResponse{FileData: fileData}, nil
}

func StartServer(ctx context.Context) {
	listener, listenError := net.Listen("tcp", ":8081")
	if listenError != nil {
		panic(listenError)
	}

	server := grpc.NewServer()
	reflection.Register(server) //Enabled for clients that support reflection

  grpc_health_v1.RegisterHealthServer(server, health.NewServer())

	proto.RegisterGobreServer(server, GobreServer{})

	//Start the server in a separate goroutine
	go func() {
		<-ctx.Done()
		server.Stop()
	}()

	//Serve the server
  serverError := server.Serve(listener)

  if serverError != nil {
    panic(serverError)
  }
}
