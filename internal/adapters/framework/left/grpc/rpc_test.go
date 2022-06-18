package rpc

import (
	"context"
	"log"
	"net"
	"os"
	"testing"

	"github.com/hex/internal/adapters/app/api"
	"github.com/hex/internal/adapters/core/airthmatic"
	"github.com/hex/internal/adapters/framework/left/grpc/pb"
	"github.com/hex/internal/adapters/framework/right/db"
	"github.com/hex/internal/ports"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc"
	"google.golang.org/grpc/test/bufconn"
)

const buffSize = 1024 * 1024

var lis *bufconn.Listener

func init() {
	var err error
	lis = bufconn.Listen(buffSize)
	grpcServer := grpc.NewServer()

	//ports
	var core ports.AirthmaticPort
	var dbasAdapter ports.DBPort
	var appAdapter ports.APIPort
	var grpcAdapter ports.GRPCPort

	dbaseDriver := os.Getenv("DB_DRIVER")
	dataSourceName := os.Getenv("DS_NAME")

	dbasAdapter, err = db.NewAdapter(dbaseDriver, dataSourceName)
	if err != nil {
		log.Fatalf("failed to initialize db connection:%v", err)
	}

	core = airthmatic.NewAdapter()

	appAdapter = api.NewAdapter(core, dbasAdapter)
	grpcAdapter = NewAdapter(appAdapter)

	pb.RegisterAirthmaticServiceServer(grpcServer, grpcAdapter)

	go func() {
		if err := grpcServer.Serve(lis); err != nil {
			log.Fatalf("test server start error:%v", err)
		}
	}()

}

func bufDialer(context.Context, string) (net.Conn, error) {
	return lis.Dial()
}

func getGrpcConnection(ctx context.Context, t *testing.T) *grpc.ClientConn {
	conn, err := grpc.DialContext(ctx, "bufnet", grpc.WithContextDialer(bufDialer), grpc.WithInsecure())
	if err != nil {
		t.Fatalf("test server start error:%v", err)
	}
	return conn
}

func TestAddition(t *testing.T) {
	ctx := context.Background()
	conn := getGrpcConnection(ctx, t)
	defer conn.Close()

	client := pb.NewAirthmaticServiceClient(conn)

	params := &pb.OperationParameter{
		A: 1,
		B: 1,
	}

	answer, err := client.GetAddition(ctx, params)

	if err != nil {
		t.Fatalf("expected:%v, got:%v", nil, err)
	}
	require.Equal(t, answer.Value, int32(2))
}

func TestSubstration(t *testing.T) {
	ctx := context.Background()
	conn := getGrpcConnection(ctx, t)
	defer conn.Close()

	client := pb.NewAirthmaticServiceClient(conn)

	params := &pb.OperationParameter{
		A: 1,
		B: 1,
	}

	answer, err := client.GetSubstraction(ctx, params)

	if err != nil {
		t.Fatalf("expected:%v, got:%v", nil, err)
	}
	require.Equal(t, answer.Value, int32(0))
}

func TestMultiplication(t *testing.T) {
	ctx := context.Background()
	conn := getGrpcConnection(ctx, t)
	defer conn.Close()

	client := pb.NewAirthmaticServiceClient(conn)

	params := &pb.OperationParameter{
		A: 1,
		B: 1,
	}

	answer, err := client.GetMultiplication(ctx, params)

	if err != nil {
		t.Fatalf("expected:%v, got:%v", nil, err)
	}
	require.Equal(t, answer.Value, int32(1))
}

func TestDivison(t *testing.T) {
	ctx := context.Background()
	conn := getGrpcConnection(ctx, t)
	defer conn.Close()

	client := pb.NewAirthmaticServiceClient(conn)

	params := &pb.OperationParameter{
		A: 1,
		B: 1,
	}

	answer, err := client.GetDivison(ctx, params)

	if err != nil {
		t.Fatalf("expected:%v, got:%v", nil, err)
	}
	require.Equal(t, answer.Value, int32(1))
}
