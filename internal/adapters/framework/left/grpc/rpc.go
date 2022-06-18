package rpc

import (
	"context"
	"fmt"

	"github.com/hex/internal/adapters/framework/left/grpc/pb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (grpca Adapter) GetAddition(ctx context.Context, req *pb.OperationParameter) (*pb.Answer, error) {
	var err error
	ans := &pb.Answer{}

	if req.GetA() == 0 || req.GetB() == 0 {
		return ans, status.Error(codes.InvalidArgument, "missign required")
	}
	answer, err := grpca.api.GetAddition(req.A, req.B)

	if err != nil {
		return ans, status.Error(codes.Internal, "unexpected error")
	}
	ans = &pb.Answer{Value: answer}
	return ans, err
}

func (grpca Adapter) GetSubstraction(ctx context.Context, req *pb.OperationParameter) (*pb.Answer, error) {
	var err error
	ans := &pb.Answer{}

	if req.GetA() == 0 || req.GetB() == 0 {
		return ans, status.Error(codes.InvalidArgument, "missign required")
	}
	answer, err := grpca.api.GetSubstraction(req.A, req.B)

	if err != nil {
		return ans, status.Error(codes.Internal, fmt.Sprintf("%v", err))
	}
	ans = &pb.Answer{Value: answer}
	return ans, err
}

func (grpca Adapter) GetMultiplication(ctx context.Context, req *pb.OperationParameter) (*pb.Answer, error) {
	var err error
	ans := &pb.Answer{}

	if req.GetA() == 0 || req.GetB() == 0 {
		return ans, status.Error(codes.InvalidArgument, "missign required")
	}
	answer, err := grpca.api.GetMultiplication(req.A, req.B)

	if err != nil {
		return ans, status.Error(codes.Internal, "unexpected error")
	}
	ans = &pb.Answer{Value: answer}
	return ans, err
}

func (grpca Adapter) GetDivison(ctx context.Context, req *pb.OperationParameter) (*pb.Answer, error) {
	var err error
	ans := &pb.Answer{}

	if req.GetA() == 0 || req.GetB() == 0 {
		return ans, status.Error(codes.InvalidArgument, "missign required")
	}
	answer, err := grpca.api.GetDivison(req.A, req.B)

	if err != nil {
		return ans, status.Error(codes.Internal, "unexpected error")
	}
	ans = &pb.Answer{Value: answer}
	return ans, err
}
