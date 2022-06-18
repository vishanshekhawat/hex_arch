package rpc

import (
	"log"
	"net"

	"github.com/hex/internal/adapters/framework/left/grpc/pb"
	"github.com/hex/internal/ports"
	"google.golang.org/grpc"
)

type Adapter struct {
	api ports.APIPort
}

func NewAdapter(api ports.APIPort) *Adapter {
	return &Adapter{api: api}
}

func (grpca Adapter) Run() {
	var err error
	listen, err := net.Listen("tcp", "9000")
	if err != nil {
		log.Fatalf("listen on port 9000: %v", err)
	}

	airthmaticServiceServer := grpca
	grpcServer := grpc.NewServer()
	pb.RegisterAirthmaticServiceServer(grpcServer, airthmaticServiceServer)

	if err = grpcServer.Serve(listen); err != nil {
		log.Fatalf("failed to serve grpc on port 9000: %v", err)
	}
}
