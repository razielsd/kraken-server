package main

import (
	"context"
	"net"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/razielsd/kraken-server/internal/app/projectservice"
	"github.com/razielsd/kraken-server/pkg/papi"
	"golang.org/x/sync/errgroup"
	"google.golang.org/grpc"

	"github.com/tmc/grpc-websocket-proxy/wsproxy"
)

func main() {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	lis, err := net.Listen("tcp", ":8842")
	if err != nil {
		return
	}
	grpcServer := grpc.NewServer()

	service := projectservice.NewProjectService()
	papi.RegisterProjectServer(grpcServer, service)

	var group errgroup.Group

	group.Go(func() error {
		return grpcServer.Serve(lis)
	})

	mux := runtime.NewServeMux(runtime.WithMarshalerOption(runtime.MIMEWildcard, &runtime.JSONPb{OrigName: true, EmitDefaults: true}))
	opts := []grpc.DialOption{
		grpc.WithInsecure(),
		grpc.WithDefaultCallOptions(grpc.MaxCallRecvMsgSize(50000000)),
	}

	group.Go(func() error {
		return papi.RegisterProjectHandlerFromEndpoint(ctx, mux, ":8842", opts)
	})

	group.Go(func() error {
		return http.ListenAndServe(":8843", wsproxy.WebsocketProxy(mux))
	})

	group.Wait()
}
