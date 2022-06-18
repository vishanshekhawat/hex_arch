package ports

import (
	"context"

	"github.com/hex/internal/adapters/framework/left/grpc/pb"
)

type GRPCPort interface {
	Run()
	GetAddition(ctx context.Context, req *pb.OperationParameter) (*pb.Answer, error)
	GetSubstraction(ctx context.Context, req *pb.OperationParameter) (*pb.Answer, error)
	GetMultiplication(ctx context.Context, req *pb.OperationParameter) (*pb.Answer, error)
	GetDivison(ctx context.Context, req *pb.OperationParameter) (*pb.Answer, error)
}
