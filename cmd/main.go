package main

import (
	"log"
	"os"

	"github.com/hex/internal/adapters/app/api"
	"github.com/hex/internal/adapters/core/airthmatic"
	gRPC "github.com/hex/internal/adapters/framework/left/grpc"
	"github.com/hex/internal/adapters/framework/right/db"
	"github.com/hex/internal/ports"
)

func main() {

	// ports
	var core ports.AirthmaticPort
	var dbasAdapter ports.DBPort
	var appAdapter ports.APIPort
	var grpcAdapter ports.GRPCPort

	dbaseDriver := os.Getenv("DB_DRIVER")
	dataSourceName := os.Getenv("DS_NAME")

	dbasAdapter, err := db.NewAdapter(dbaseDriver, dataSourceName)
	if err != nil {
		log.Fatalf("failed to initialize db connection:%v", err)
	}
	defer dbasAdapter.CloseDbConnection()
	core = airthmatic.NewAdapter()

	appAdapter = api.NewAdapter(core, dbasAdapter)
	grpcAdapter = gRPC.NewAdapter(appAdapter)
	grpcAdapter.Run()

}
