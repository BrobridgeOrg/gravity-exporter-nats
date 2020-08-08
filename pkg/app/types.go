package app

import (
	"github.com/BrobridgeOrg/gravity-exporter-nats/pkg/grpc_server"
	"github.com/BrobridgeOrg/gravity-exporter-nats/pkg/mux_manager"
)

type App interface {
	GetGRPCServer() grpc_server.Server
	GetMuxManager() mux_manager.Manager
}
