package server

import (
	"context"
	"testing"

	proto "main/proto/gobre"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc"
)

func TestStartServer(t *testing.T){
	ctx, cancel := context.WithCancel(context.Background())
	go StartServer(ctx)

	conn, connErr := grpc.NewClient(
		":8081", 
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if connErr != nil {
		t.Fatal(connErr)
	}
	defer conn.Close()

	client := proto.NewGobreClient(conn)

	request := &proto.FileRequest{
		FileData: []byte("<html><b>test</b></html>"),
		NewFileType: "pdf",
		OriginalFileType: "html",
	}
	response, responseErr := client.HandleFileRequest(context.Background(), request)
	
	if responseErr != nil {
		t.Fatal(responseErr)
	}
	
  if string(response.GetFileData()) == "" {
		t.Errorf("No file data in the response.")
	}

	cancel()
}
