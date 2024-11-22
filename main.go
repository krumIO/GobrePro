package main

import (
	"context"
	server "main/server"
)

func main(){
	ctx := context.Background()
	server.StartServer(ctx)
}
